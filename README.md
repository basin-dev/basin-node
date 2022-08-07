<!--- ![](docs/green.png) -->
<!--- https://shields.io/ (once the repo is public) -->
<!--- ![](docs/basin.png) -->

```go
______  ___   _____ _____ _   _ 
| ___ \/ _ \ /  ___|_   _| \ | |
| |_/ / /_\ \\ '--.  | | |  \| |
| ___ \  _  | '--. \ | | | . ' |
| |_/ / | | |/\__/ /_| |_| |\  |
\____/\_| |_/\____/ \___/\_| \_/
```

# Basin node

## Development

### Getting started

Clone the `basin-node` repo:
```
git@github.com:basin-dev/basin-node.git
```

Enter the `src` directory:
```
cd src
```

### Command Line Interface (CLI)

#### Background

[Cobra](https://github.com/spf13/cobra) is used to build the CLI for the Basin Node app

[Cobra-CLI generator](https://github.com/spf13/cobra-cli/blob/main/README.md) is used to bootstrap application scaffolding for rapid development

[Viper](https://github.com/spf13/viper) is used as a registry for all future application configuration needs as a 12 factor app

[OpenAPI Generator](https://openapi-generator.tech/) is used to automatically generate server stubs and an API client for the node's HTTP interface.

#### Adding a new command

Use the Cobra-CLI generator to add a new command:
```
cobra-cli add [COMMAND_NAME]
```
