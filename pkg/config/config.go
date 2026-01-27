package config

//go:generate go run ./gen

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	UsernameField = field.StringField(
		"username",
		field.WithDescription("The username of the CouchDB admin account"),
		field.WithRequired(true),
	)
	PasswordField = field.StringField(
		"password",
		field.WithDescription("The password of the CouchDB admin account"),
		field.WithRequired(true),
	)
	InstanceHostField = field.StringField(
		"instance-url",
		field.WithDescription("The url to the CouchDB instance. Include :port if needed"),
		field.WithRequired(true),
	)

	// ConfigurationFields defines the external configuration required for the
	// connector to run.
	ConfigurationFields = []field.SchemaField{UsernameField, PasswordField, InstanceHostField}

	// FieldRelationships defines relationships between the fields.
	FieldRelationships = []field.SchemaFieldRelationship{}

	// Config is the configuration schema for the connector.
	Config = field.Configuration{
		Fields:      ConfigurationFields,
		Constraints: FieldRelationships,
	}
)
