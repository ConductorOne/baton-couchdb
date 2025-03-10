package main

import (
	"fmt"

	"github.com/conductorone/baton-sdk/pkg/field"
	"github.com/spf13/viper"
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
	InstanceHost = field.StringField(
		"instance-url",
		field.WithDescription("The url to the CouchDB instance. Include :port if needed"),
		field.WithRequired(true),
	)
)

var (
	// ConfigurationFields defines the external configuration required for the
	// connector to run. Note: these fields can be marked as optional or
	// required.
	ConfigurationFields = []field.SchemaField{UsernameField, PasswordField, InstanceHost}

	// FieldRelationships defines relationships between the fields listed in
	// ConfigurationFields that can be automatically validated. For example, a
	// username and password can be required together, or an access token can be
	// marked as mutually exclusive from the username password pair.
	FieldRelationships = []field.SchemaFieldRelationship{}
)

// ValidateConfig is run after the configuration is loaded, and should return an
// error if it isn't valid. Implementing this function is optional, it only
// needs to perform extra validations that cannot be encoded with configuration
// parameters.
func ValidateConfig(v *viper.Viper) error {
	username := v.GetString(UsernameField.FieldName)
	password := v.GetString(PasswordField.FieldName)
	instanceURL := v.GetString(InstanceHost.FieldName)

	if username == "" || password == "" || instanceURL == "" {
		return fmt.Errorf("the required fields '--%s', '--%s' and '--%s' can't be empty", UsernameField.FieldName, PasswordField.FieldName, InstanceHost.FieldName)
	}

	return nil
}
