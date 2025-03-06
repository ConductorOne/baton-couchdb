package connector

import (
	"context"

	"github.com/conductorone/baton-couchdb/pkg/client"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	resourceSdk "github.com/conductorone/baton-sdk/pkg/types/resource"
)

type databaseBuilder struct {
	Client *client.CouchDBClient
}

func (b *databaseBuilder) ResourceType(_ context.Context) *v2.ResourceType {
	return databaseResourceType
}

func (b *databaseBuilder) List(ctx context.Context, _ *v2.ResourceId, _ *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	var databaseResources []*v2.Resource

	dbs, err := b.Client.ListAllDataBases(ctx)
	if err != nil {
		return nil, "", nil, err
	}

	for _, dbName := range dbs {
		dbResource, err := parseIntoDatabaseResource(dbName)
		if err != nil {
			return nil, "", nil, err
		}
		databaseResources = append(databaseResources, dbResource)
	}

	return databaseResources, "", nil, nil
}

func (b *databaseBuilder) Entitlements(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func (b *databaseBuilder) Grants(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func parseIntoDatabaseResource(database string) (*v2.Resource, error) {
	profile := map[string]interface{}{
		"database_name": database,
	}

	groupTraitOptions := []resourceSdk.GroupTraitOption{
		resourceSdk.WithGroupProfile(profile),
	}

	ret, err := resourceSdk.NewGroupResource(
		database,
		databaseResourceType,
		database,
		groupTraitOptions,
		resourceSdk.WithAnnotation(
			&v2.ChildResourceType{ResourceTypeId: userResourceType.Id},
			&v2.ChildResourceType{ResourceTypeId: roleResourceType.Id},
		),
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func newDatabaseBuilder(c *client.CouchDBClient) *databaseBuilder {
	return &databaseBuilder{
		Client: c,
	}
}
