package main

import (
	cfg "github.com/conductorone/baton-couchdb/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("couchdb", cfg.Config)
}
