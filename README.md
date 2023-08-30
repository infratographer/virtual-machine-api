![logo](https://github.com/infratographer/website/blob/main/source/theme/assets/pictures/logo.jpg?raw=true)
# virtual-machine-api - The virtual machine management API

`virtual-machine-api` is a GraphQL API for managing virtual machines. It is written in GO and uses the [ent](https://entgo.io) and [gqlgen](https://gqlgen.com/) as framework for defining an abstract model of a virtual machine for the Infratographer ecosystem.

The API doesn't create virtual machines but rather provides a way to manage the state of virtual machines. It is up to a client referred to as a [resource provider](https://infratographer.com/docs/architecture/foundational-resources/#resource-providers) to create the virtual machine in the underlying infrastructure and then use the API to manage the state of the virtual machine.

## Getting Started

This project uses [`devcontainers`](https://containers.dev) to provide a development environment. To get started, you will need to install [Docker](https://www.docker.com/) and [VSCode](https://code.visualstudio.com/). Once you have those installed, you can open the project in VSCode and click the "Reopen in Container" button in the bottom right corner of the window. This will build the devcontainer and open the project in a container with all the dependencies installed.

Once the devcontainer is built, you can run the following commands:

```bash
make go-run
```

This will start the API server and VSCode will [forward the port](https://code.visualstudio.com/docs/remote/ssh#_forwarding-a-port-creating-ssh-tunnel). You can then navigate to http://localhost:7911/playground to view the GraphQL playground.