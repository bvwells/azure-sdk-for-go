// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"context"
	"os"
	"testing"
)

type emulatorTests struct {
	host string
	key  string
}

func newEmulatorTests(t *testing.T) *emulatorTests {
	envCheck := os.Getenv("EMULATOR")
	if envCheck == "" {
		t.Skip("set EMULATOR environment variable to run this test")
	}

	return &emulatorTests{
		host: "https://localhost:8081/",
		key:  "C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw==",
	}
}

func (e *emulatorTests) getClient(t *testing.T) *CosmosClient {
	cred, _ := NewSharedKeyCredential(e.key)
	client, err := NewClientWithSharedKey(e.host, cred, nil)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	return client
}

func (e *emulatorTests) createDatabase(
	t *testing.T,
	ctx context.Context,
	client *CosmosClient,
	dbName string) *Database {
	database := DatabaseProperties{Id: dbName}
	resp, err := client.CreateDatabase(ctx, database, nil, nil)
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}

	if resp.DatabaseProperties.Id != database.Id {
		t.Errorf("Unexpected id match: %v", resp.DatabaseProperties)
	}

	return resp.DatabaseProperties.Database
}

func (e *emulatorTests) deleteDatabase(
	t *testing.T,
	ctx context.Context,
	database *Database) {
	_, err := database.Delete(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to delete database: %v", err)
	}
}
