package connector

import (
	"context"
	"fmt"
	"strconv"

	"github.com/conductorone/baton-couchdb/pkg/client"
	resourceSdk "github.com/conductorone/baton-sdk/pkg/types/resource"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
)

type userBuilder struct {
	Client *client.CouchDBClient
}

func (b *userBuilder) ResourceType(_ context.Context) *v2.ResourceType {
	return userResourceType
}

// List returns all the users from the database as resource objects.
// Users include a UserTrait because they are the 'shape' of a standard user.
func (b *userBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, _ *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	if parentResourceID == nil {
		return nil, "", nil, nil
	}

	var users []client.User
	var userResources []*v2.Resource

	dbName := parentResourceID.Resource
	dbSecurityObject, err := b.Client.GetSecurityObject(ctx, dbName)
	if err != nil {
		return nil, "", nil, err
	}

	if dbSecurityObject == nil {
		return nil, "", nil, fmt.Errorf("the security object of '%s' database couldn't be retrieved", dbName)
	}

	adminUsers, err := extractUsers(dbSecurityObject.Admins, dbName)
	if err != nil {
		return nil, "", nil, err
	}
	users = append(users, adminUsers...)

	memberUsers, err := extractUsers(dbSecurityObject.Members, dbName)
	if err != nil {
		return nil, "", nil, err
	}
	users = append(users, memberUsers...)

	for _, user := range users {
		userResource, err := parseIntoUserResource(user, parentResourceID)
		if err != nil {
			return nil, "", nil, err
		}
		userResources = append(userResources, userResource)
	}

	return userResources, "", nil, nil
}

func extractUsers(secComponent client.SecurityComponent, dbName string) ([]client.User, error) {
	if len(secComponent.Names) != len(secComponent.Roles) {
		return nil, nil // fmt.Errorf("the amount of names and roles don't match. It's not possible to identify user roles")
	}

	var users []client.User
	for i, name := range secComponent.Names {
		users = append(
			users,
			client.User{
				Username: name,
				Role:     strconv.Itoa(i),
				Database: dbName,
			},
		)
	}

	for i, role := range secComponent.Roles {
		if len(users) > i && users[i].Role == strconv.Itoa(i) {
			users[i].Role = role
		}
	}

	return users, nil
}

// Entitlements always returns an empty slice for users.
func (b *userBuilder) Entitlements(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

// Grants always returns an empty slice for users since they don't have any entitlements.
func (b *userBuilder) Grants(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func parseIntoUserResource(user client.User, parentResourceID *v2.ResourceId) (*v2.Resource, error) {
	var userStatus = v2.UserTrait_Status_STATUS_ENABLED
	id := createID(user.Database, user.Username)
	displayName := "db:" + user.Database + " / " + user.Username

	profile := map[string]interface{}{
		"id":       id,
		"username": user.Username,
		"database": user.Database,
	}

	userTraits := []resourceSdk.UserTraitOption{
		resourceSdk.WithUserProfile(profile),
		resourceSdk.WithStatus(userStatus),
	}

	return resourceSdk.NewUserResource(
		displayName,
		userResourceType,
		id,
		userTraits,
		resourceSdk.WithParentResourceID(parentResourceID),
	)
}

func newUserBuilder(c *client.CouchDBClient) *userBuilder {
	return &userBuilder{
		Client: c,
	}
}
