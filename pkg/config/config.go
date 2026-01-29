package config

//go:generate go run ./gen

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var Config = field.NewConfiguration([]field.SchemaField{
	field.StringField(
		"username",
		field.WithDescription("The username of the CouchDB admin account"),
		field.WithRequired(true),
	),
	field.StringField(
		"password",
		field.WithDescription("The password of the CouchDB admin account"),
		field.WithRequired(true),
	),
	field.StringField(
		"instance-url",
		field.WithDescription("The url to the CouchDB instance. Include :port if needed"),
		field.WithRequired(true),
	),
})

func ValidateConfig(c *Couchdb) error {
	return nil
}
