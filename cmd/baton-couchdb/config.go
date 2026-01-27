package main

import (
	"fmt"

	cfg "github.com/conductorone/baton-couchdb/pkg/config"
	"github.com/spf13/viper"
)

// ValidateConfig is run after the configuration is loaded, and should return an
// error if it isn't valid. Implementing this function is optional, it only
// needs to perform extra validations that cannot be encoded with configuration
// parameters.
func ValidateConfig(v *viper.Viper) error {
	username := v.GetString(cfg.UsernameField.FieldName)
	password := v.GetString(cfg.PasswordField.FieldName)
	instanceURL := v.GetString(cfg.InstanceHostField.FieldName)

	if username == "" || password == "" || instanceURL == "" {
		return fmt.Errorf("the required fields '--%s', '--%s' and '--%s' can't be empty", cfg.UsernameField.FieldName, cfg.PasswordField.FieldName, cfg.InstanceHostField.FieldName)
	}

	return nil
}
