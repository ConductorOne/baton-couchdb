package connector

import (
	"context"
	"io"

	"github.com/conductorone/baton-couchdb/pkg/client"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
)

type Connector struct {
	Client *client.CouchDBClient
}

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced from the upstream service.
func (c *Connector) ResourceSyncers(_ context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{
		newUserBuilder(c.Client),
		newRoleBuilder(c.Client),
		newDatabaseBuilder(c.Client),
	}
}

// Asset takes an input AssetRef and attempts to fetch it using the connector's authenticated http client
// It streams a response, always starting with a metadata object, following by chunked payloads for the asset.
func (d *Connector) Asset(_ context.Context, _ *v2.AssetRef) (string, io.ReadCloser, error) {
	return "", nil, nil
}

// Metadata returns metadata about the connector.
func (d *Connector) Metadata(_ context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "CouchDB Connector",
		Description: "Baton connector made to work with a CouchDB instance. It can retrieve Users and Roles information of the databases on the instance.",
	}, nil
}

// Validate is called to ensure that the connector is properly configured. It should exercise any API credentials
// to be sure that they are valid.
func (d *Connector) Validate(_ context.Context) (annotations.Annotations, error) {
	return nil, nil
}

// New returns a new instance of the connector.
func New(ctx context.Context, username, password, instanceURL string) (*Connector, error) {
	c, err := client.New(
		ctx,
		client.WithBasicAuth(username, password),
		client.WithInstanceURL(instanceURL),
	)
	if err != nil {
		return nil, err
	}

	return &Connector{
		Client: c,
	}, nil
}
