package connector

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-couchdb/pkg/client"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	"github.com/conductorone/baton-sdk/pkg/types/entitlement"
	"github.com/conductorone/baton-sdk/pkg/types/grant"
	resourceSdk "github.com/conductorone/baton-sdk/pkg/types/resource"
)

const permissionName = "assigned"

type roleBuilder struct {
	Client     *client.CouchDBClient
	UsersCache []client.User
}

func (b *roleBuilder) ResourceType(_ context.Context) *v2.ResourceType {
	return roleResourceType
}

func (b *roleBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, _ *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	if parentResourceID == nil {
		return nil, "", nil, nil
	}

	var roleResources []*v2.Resource
	dbName := parentResourceID.Resource

	users, err := b.retrieveDBUsers(ctx, dbName)
	if err != nil {
		return nil, "", nil, err
	}

	for _, user := range users {
		roleResource, err := parseIntoRoleResource(user, parentResourceID)
		if err != nil {
			return nil, "", nil, err
		}
		roleResources = append(roleResources, roleResource)
	}

	return roleResources, "", nil, nil
}

func (b *roleBuilder) retrieveDBUsers(ctx context.Context, dbName string) ([]client.User, error) {
	var users []client.User

	dbSecurityObject, err := b.Client.GetSecurityObject(ctx, dbName)
	if err != nil {
		return nil, err
	}

	if dbSecurityObject == nil {
		return nil, fmt.Errorf("the security object of '%s' database couldn't be retrieved", dbName)
	}

	adminUsers, err := extractUsers(dbSecurityObject.Admins, dbName)
	if err != nil {
		return nil, err
	}
	users = append(users, adminUsers...)

	memberUsers, err := extractUsers(dbSecurityObject.Members, dbName)
	if err != nil {
		return nil, err
	}
	users = append(users, memberUsers...)

	b.UsersCache = append(b.UsersCache, users...)

	return users, nil
}

func (b *roleBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var ret []*v2.Entitlement

	assigmentOptions := []entitlement.EntitlementOption{
		entitlement.WithGrantableTo(userResourceType),
		entitlement.WithDescription(resource.Description),
		entitlement.WithDisplayName(resource.DisplayName),
	}
	ret = append(ret, entitlement.NewPermissionEntitlement(resource, permissionName, assigmentOptions...))

	return ret, "", nil, nil
}

func (b *roleBuilder) Grants(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	var ret []*v2.Grant

	for _, user := range b.UsersCache {
		userRole := createID(user.Database, user.Role)
		if userRole == resource.Id.Resource {
			userResource, err := parseIntoUserResource(user, nil)
			if err != nil {
				return nil, "", nil, err
			}

			membershipGrant := grant.NewGrant(resource, permissionName, userResource.Id)
			ret = append(ret, membershipGrant)
		}
	}

	return ret, "", nil, nil
}

func parseIntoRoleResource(user client.User, parentResourceID *v2.ResourceId) (*v2.Resource, error) {
	id := createID(user.Database, user.Role)
	displayName := "db:" + user.Database + " / " + user.Role

	profile := map[string]interface{}{
		"id":       id,
		"username": user.Role,
		"database": user.Database,
	}

	roleTraits := []resourceSdk.RoleTraitOption{
		resourceSdk.WithRoleProfile(profile),
	}

	return resourceSdk.NewRoleResource(
		displayName,
		roleResourceType,
		id,
		roleTraits,
		resourceSdk.WithParentResourceID(parentResourceID),
	)
}

func newRoleBuilder(c *client.CouchDBClient) *roleBuilder {
	return &roleBuilder{
		Client:     c,
		UsersCache: []client.User{},
	}
}

func createID(dbName, resourceName string) string {
	return dbName + "/" + resourceName
}
