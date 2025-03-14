![Baton Logo](./baton-logo.png)

# `baton-couchdb` [![Go Reference](https://pkg.go.dev/badge/github.com/conductorone/baton-couchdb.svg)](https://pkg.go.dev/github.com/conductorone/baton-couchdb) ![main ci](https://github.com/conductorone/baton-couchdb/actions/workflows/main.yaml/badge.svg)

`baton-couchdb` is a connector for [CouchDB](https://couchdb.apache.org/) built using the [Baton SDK](https://github.com/conductorone/baton-sdk).
This connector requires the `username`, `password` and `instance-url` args. When synced, the names of the users and roles of each database of the instance will be listed.
Granting and revoking of roles is not supported.

Check out [Baton](https://github.com/conductorone/baton) to learn more the project in general.

# Getting Started

## brew

```
brew install conductorone/baton/baton conductorone/baton/baton-couchdb
baton-couchdb
baton resources
```

## docker

```
docker run --rm -v $(pwd):/out -e BATON_DOMAIN_URL=domain_url -e BATON_API_KEY=apiKey -e BATON_USERNAME=username ghcr.io/conductorone/baton-couchdb:latest -f "/out/sync.c1z"
docker run --rm -v $(pwd):/out ghcr.io/conductorone/baton:latest -f "/out/sync.c1z" resources
```

## source

```
go install github.com/conductorone/baton/cmd/baton@main
go install github.com/conductorone/baton-couchdb/cmd/baton-couchdb@main

baton-couchdb

baton resources
```

# Data Model

`baton-couchdb` will pull down information about the following resources:
- Users
- Roles

# Contributing, Support and Issues

We started Baton because we were tired of taking screenshots and manually
building spreadsheets. We welcome contributions, and ideas, no matter how
small&mdash;our goal is to make identity and permissions sprawl less painful for
everyone. If you have questions, problems, or ideas: Please open a GitHub Issue!

See [CONTRIBUTING.md](https://github.com/ConductorOne/baton/blob/main/CONTRIBUTING.md) for more details.

# `baton-couchdb` Command Line Usage

```
baton-couchdb

Usage:
  baton-couchdb [flags]
  baton-couchdb [command]

Available Commands:
  capabilities       Get connector capabilities
  completion         Generate the autocompletion script for the specified shell
  help               Help about any command

Flags:
      --client-id string             The client ID used to authenticate with ConductorOne ($BATON_CLIENT_ID)
      --client-secret string         The client secret used to authenticate with ConductorOne ($BATON_CLIENT_SECRET)
  -f, --file string                  The path to the c1z file to sync with ($BATON_FILE) (default "sync.c1z")
  -h, --help                         help for baton-couchdb
      --log-format string            The output format for logs: json, console ($BATON_LOG_FORMAT) (default "json")
      --log-level string             The log level: debug, info, warn, error ($BATON_LOG_LEVEL) (default "info")
  -p, --provisioning                 If this connector supports provisioning, this must be set in order for provisioning actions to be enabled ($BATON_PROVISIONING)
      --ticketing                    This must be set to enable ticketing support ($BATON_TICKETING)
  -v, --version                      version for baton-couchdb

  --username                         The username of the CouchDB admin account
  --password                         The password of the CouchDB admin account
  --instance-url                     The url to the CouchDB instance. Include :port if needed

Use "baton-couchdb [command] --help" for more information about a command.
```
