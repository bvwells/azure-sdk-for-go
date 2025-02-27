// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"context"
	"errors"
)

// A Database lets you perform read, update, change throughput, and delete database operations.
type Database struct {
	// The Id of the Cosmos database
	Id string
	// The client associated with the Cosmos database
	client *CosmosClient
	// The resource link
	link string
}

func newDatabase(id string, client *CosmosClient) *Database {
	return &Database{
		Id:     id,
		client: client,
		link:   createLink("", pathSegmentDatabase, id)}
}

// GetContainer returns a Container object for the container.
// id - The id of the container.
func (db *Database) GetContainer(id string) (*Container, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	return newContainer(id, db), nil
}

// CreateContainer creates a container in the Cosmos database.
// ctx - The context for the request.
// containerProperties - The properties for the container.
// throughputProperties - Optional throughput configuration of the container
// requestOptions - Optional parameters for the request.
func (db *Database) CreateContainer(
	ctx context.Context,
	containerProperties ContainerProperties,
	throughputProperties *ThroughputProperties,
	requestOptions *ContainerRequestOptions) (ContainerResponse, error) {
	if requestOptions == nil {
		requestOptions = &ContainerRequestOptions{}
	}

	operationContext := cosmosOperationContext{
		resourceType:    resourceTypeCollection,
		resourceAddress: db.link,
	}

	path, err := generatePathForNameBased(resourceTypeCollection, db.link, true)
	if err != nil {
		return ContainerResponse{}, err
	}

	container, err := db.GetContainer(containerProperties.Id)
	if err != nil {
		return ContainerResponse{}, err
	}

	azResponse, err := db.client.connection.sendPostRequest(
		path,
		ctx,
		containerProperties,
		operationContext,
		requestOptions,
		throughputProperties.addHeadersToRequest)
	if err != nil {
		return ContainerResponse{}, err
	}

	return newContainerResponse(azResponse, container)
}

// Read obtains the information for a Cosmos database.
// ctx - The context for the request.
// requestOptions - Optional parameters for the request.
func (db *Database) Read(
	ctx context.Context,
	requestOptions *DatabaseRequestOptions) (DatabaseResponse, error) {
	if requestOptions == nil {
		requestOptions = &DatabaseRequestOptions{}
	}

	operationContext := cosmosOperationContext{
		resourceType:    resourceTypeDatabase,
		resourceAddress: db.link,
	}

	path, err := generatePathForNameBased(resourceTypeDatabase, db.link, false)
	if err != nil {
		return DatabaseResponse{}, err
	}

	azResponse, err := db.client.connection.sendGetRequest(
		path,
		ctx,
		operationContext,
		requestOptions,
		nil)
	if err != nil {
		return DatabaseResponse{}, err
	}

	return newDatabaseResponse(azResponse, db)
}

// ReadThroughput obtains the provisioned throughput information for the database.
// ctx - The context for the request.
// requestOptions - Optional parameters for the request.
func (db *Database) ReadThroughput(
	ctx context.Context,
	requestOptions *ThroughputRequestOptions) (ThroughputResponse, error) {
	if requestOptions == nil {
		requestOptions = &ThroughputRequestOptions{}
	}

	rid, err := db.getRID(ctx)
	if err != nil {
		return ThroughputResponse{}, err
	}

	offers := &cosmosOffers{connection: db.client.connection}
	return offers.ReadThroughputIfExists(ctx, rid, requestOptions)
}

// ReplaceThroughput updates the provisioned throughput for the database.
// ctx - The context for the request.
// throughputProperties - The throughput configuration of the database.
// requestOptions - Optional parameters for the request.
func (db *Database) ReplaceThroughput(
	ctx context.Context,
	throughputProperties ThroughputProperties,
	requestOptions *ThroughputRequestOptions) (ThroughputResponse, error) {
	if requestOptions == nil {
		requestOptions = &ThroughputRequestOptions{}
	}

	rid, err := db.getRID(ctx)
	if err != nil {
		return ThroughputResponse{}, err
	}

	offers := &cosmosOffers{connection: db.client.connection}
	return offers.ReadThroughputIfExists(ctx, rid, requestOptions)
}

// Delete a Cosmos database.
// ctx - The context for the request.
// requestOptions - Optional parameters for the request.
func (db *Database) Delete(
	ctx context.Context,
	requestOptions *DatabaseRequestOptions) (DatabaseResponse, error) {
	if requestOptions == nil {
		requestOptions = &DatabaseRequestOptions{}
	}

	operationContext := cosmosOperationContext{
		resourceType:    resourceTypeDatabase,
		resourceAddress: db.link,
	}

	path, err := generatePathForNameBased(resourceTypeDatabase, db.link, false)
	if err != nil {
		return DatabaseResponse{}, err
	}

	azResponse, err := db.client.connection.sendDeleteRequest(
		path,
		ctx,
		operationContext,
		requestOptions,
		nil)
	if err != nil {
		return DatabaseResponse{}, err
	}

	return newDatabaseResponse(azResponse, db)
}

func (db *Database) getRID(ctx context.Context) (string, error) {
	dbResponse, err := db.Read(ctx, nil)
	if err != nil {
		return "", err
	}

	return dbResponse.DatabaseProperties.ResourceId, nil
}
