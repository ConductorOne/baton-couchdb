package config

//go:generate go run ./gen

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var Config = field.NewConfiguration(
	[]field.SchemaField{
		field.StringField(
			"username",
			field.WithDescription("The username of the CouchDB admin account"),
			field.WithRequired(true),
			field.WithDisplayName("Username"),
			field.WithPlaceholder("admin"),
		),
		field.StringField(
			"password",
			field.WithDescription("The password of the CouchDB admin account"),
			field.WithRequired(true),
			field.WithIsSecret(true),
			field.WithDisplayName("Password"),
			field.WithPlaceholder("password"),
		),
		field.StringField(
			"instance-url",
			field.WithDescription("The url to the CouchDB instance. Include :port if needed"),
			field.WithRequired(true),
			field.WithDisplayName("Instance URL"),
			field.WithPlaceholder("https://couchdb.example.com:5984"),
		),
	},
	field.WithConnectorDisplayName("CouchDB"),
	field.WithIconUrl("/static/app-icons/couchdb.svg"),
	field.WithHelpUrl("/docs/baton/couchdb"),
)
