// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

type managedIdentityIDKind int

const (
	miClientID   managedIdentityIDKind = 0
	miResourceID managedIdentityIDKind = 1
)

// ManagedIDKind identifies the ID of a managed identity as either a client or resource ID
type ManagedIDKind interface {
	fmt.Stringer
	idKind() managedIdentityIDKind
}

// ClientID is an identity's client ID. Use it with ManagedIdentityCredentialOptions, for example:
// ManagedIdentityCredentialOptions{ID: ClientID("7cf7db0d-...")}
type ClientID string

func (ClientID) idKind() managedIdentityIDKind {
	return miClientID
}

func (c ClientID) String() string {
	return string(c)
}

// ResourceID is an identity's resource ID. Use it with ManagedIdentityCredentialOptions, for example:
// ManagedIdentityCredentialOptions{ID: ResourceID("/subscriptions/...")}
type ResourceID string

func (ResourceID) idKind() managedIdentityIDKind {
	return miResourceID
}

func (r ResourceID) String() string {
	return string(r)
}

// ManagedIdentityCredentialOptions contains parameters that can be used to configure the pipeline used with Managed Identity Credential.
// All zero-value fields will be initialized with their default values.
type ManagedIdentityCredentialOptions struct {
	// ID is the ID of a managed identity the credential should authenticate. Set this field to use a specific identity
	// instead of the hosting environment's default. The value may be the identity's client ID or resource ID, but note that
	// some platforms don't accept resource IDs.
	ID ManagedIDKind

	// HTTPClient sets the transport for making HTTP requests.
	// Leave this as nil to use the default HTTP transport.
	HTTPClient policy.Transporter

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry policy.TelemetryOptions

	// Logging configures the built-in logging policy behavior.
	Logging policy.LogOptions
}

// ManagedIdentityCredential attempts authentication using a managed identity that has been assigned to the deployment environment. This authentication type works in several
// managed identity environments such as Azure VMs, App Service, Azure Functions, Azure CloudShell, among others. More information about configuring managed identities can be found here:
// https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview
type ManagedIdentityCredential struct {
	id     ManagedIDKind
	client *managedIdentityClient
}

// NewManagedIdentityCredential creates a credential instance capable of authenticating an Azure managed identity in any hosting environment
// supporting managed identities. See https://docs.microsoft.com/azure/active-directory/managed-identities-azure-resources/overview for more
// information about Azure Managed Identity.
// options: ManagedIdentityCredentialOptions that configure the pipeline for requests sent to Azure Active Directory.
func NewManagedIdentityCredential(options *ManagedIdentityCredentialOptions) (*ManagedIdentityCredential, error) {
	// Create a new Managed Identity Client with default options
	if options == nil {
		options = &ManagedIdentityCredentialOptions{}
	}
	client := newManagedIdentityClient(options)
	msiType, err := client.getMSIType()
	// If there is an error that means that the code is not running in a Managed Identity environment
	if err != nil {
		credErr := &CredentialUnavailableError{credentialType: "Managed Identity Credential", message: "Please make sure you are running in a managed identity environment, such as a VM, Azure Functions, Cloud Shell, etc..."}
		logCredentialError(credErr.credentialType, credErr)
		return nil, credErr
	}
	// Assign the msiType discovered onto the client
	client.msiType = msiType
	// check if no clientID is specified then check if it exists in an environment variable
	id := options.ID
	if id == nil {
		cID := os.Getenv("AZURE_CLIENT_ID")
		if cID != "" {
			id = ClientID(cID)
		} else {
			rID := os.Getenv("AZURE_RESOURCE_ID")
			if rID != "" {
				id = ResourceID(rID)
			}
		}
	}
	return &ManagedIdentityCredential{id: id, client: client}, nil
}

// GetToken obtains an AccessToken from the Managed Identity service if available.
// scopes: The list of scopes for which the token will have access.
// Returns an AccessToken which can be used to authenticate service client calls.
func (c *ManagedIdentityCredential) GetToken(ctx context.Context, opts policy.TokenRequestOptions) (*azcore.AccessToken, error) {
	if opts.Scopes == nil {
		err := &AuthenticationFailedError{msg: "must specify a resource in order to authenticate"}
		addGetTokenFailureLogs("Managed Identity Credential", err, true)
		return nil, err
	}
	if len(opts.Scopes) != 1 {
		err := &AuthenticationFailedError{msg: "can only specify one resource to authenticate with ManagedIdentityCredential"}
		addGetTokenFailureLogs("Managed Identity Credential", err, true)
		return nil, err
	}
	// managed identity endpoints require an AADv1 resource (i.e. token audience), not a v2 scope, so we remove "/.default" here
	scopes := []string{strings.TrimSuffix(opts.Scopes[0], defaultSuffix)}
	tk, err := c.client.authenticate(ctx, c.id, scopes)
	if err != nil {
		addGetTokenFailureLogs("Managed Identity Credential", err, true)
		return nil, err
	}
	logGetTokenSuccess(c, opts)
	logMSIEnv(c.client.msiType)
	return tk, err
}

var _ azcore.TokenCredential = (*ManagedIdentityCredential)(nil)
