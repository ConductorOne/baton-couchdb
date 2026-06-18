package main

import (
	"context"

	cfg "github.com/conductorone/baton-couchdb/pkg/config"
	"github.com/conductorone/baton-couchdb/pkg/connector"
	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/conductorone/baton-sdk/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/connectorrunner"
)

var version = "dev"

func main() {
	ctx := context.Background()
	config.RunConnector(
		ctx,
		"baton-couchdb",
		version,
		cfg.Config,
		getConnector,
		connectorrunner.WithDefaultCapabilitiesConnectorBuilderV2(&connector.Connector{}),
	)
}

func getConnector(ctx context.Context, cc *cfg.Couchdb, opts *cli.ConnectorOpts) (connectorbuilder.ConnectorBuilderV2, []connectorbuilder.Opt, error) {
	c, err := connector.New(ctx, cc.Username, cc.Password, cc.InstanceUrl)
	if err != nil {
		return nil, nil, err
	}
	return c, nil, nil
}
