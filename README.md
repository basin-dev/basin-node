<!--- ![](docs/green.png) -->
<!--- https://shields.io/ (once the repo is public) -->
![](docs/basin.png)

# Basin node

## Getting started

Clone the `basin-node` repo:
```
git@github.com:basin-dev/basin-node.git
```

Enter the `src` directory:
```
cd src
```

*Question: what is a key store file?*

### Add a new keystore file for the given DID

*Question: what are our instructions for determining your DID?*

Enter a did, a name for your private key, and a password after running:
```
go run . auth add
```

*Question: why are each of these necessary?*
*Question: what does it mean to be the node's default signer?*

### Extract and print the info from your keystore file

This does not work as expected:
```
go run . auth {did} {pw}
```

### Delete the keystore for the given DID

You can remove the keystore for a DID by running:
```
go run . auth forget {did}
```

### Start the Basin node

Specify your DID and password by running:
```
go run . up --did={did} --pw={pw}
```

### Use interactive CLI

Open a second tab with `src` as your current directory again

Attach to the Basin node with interactive CLI by running:
```
go run . attach --http=http://127.0.0.1:8555
```

## Producer

From within the Basin node with interactive CLI run:

Register your first resource:
```

```

## Consumer

## Development

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