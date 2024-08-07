# mygreeterv3



## Prerequisites

### Installations
- Follow the steps to install [Go](https://go.dev/doc/install) if you do not already have it.

- Follow the steps to install [Docker](https://docs.docker.com/engine/install/) if you do not already have it.

- Set up [GOPROXY for AKS](https://dev.azure.com/msazure/CloudNativeCompute/_wiki/wikis/CloudNativeCompute.wiki/190872/aks-go-proxy)



## Setup and Development

Note that we use the remote aks middleware. This middleware is responsible for features such as logging, retry, and input validation. To learn more, please visit the [repo](https://github.com/Azure/aks-middleware/tree/main).

### Initialize service

```bash
./init.sh
# Follow instructions from the scripts to create the api module, etc.
```

### Run Service Locally

There is a simple way to run the MyGreeter service, after everything has been properly generated. Inside the MyGreeter directory, you can run the client, demoserver and server.

#### Server

To run the server:

```bash
go run go.goms.io/aks/rp/mygreeterv3/server/cmd/server start 
```

By default the server starts in port `localhost:50051` and the enable-azureSDK-calls flag is set to false.

To run the server with the azureSDK calls enabled:

```bash
go run go.goms.io/aks/rp/mygreeterv3/server/cmd/server start --enable-azureSDK-calls true --subscription-id <sub_id>
```

By default, the sayHello calls are served directly by the server. In order to forward the call to the demoserver:

```bash
go run go.goms.io/aks/rp/mygreeterv3/server/cmd/server start --remote-addr <remote_addr>
```

#### Client

To run the client:

```bash
go run go.goms.io/aks/rp/mygreeterv3/server/cmd/client hello
```

By default the client sends messages to port `localhost:50051`. This can be changed by running

```bash
go run go.goms.io/aks/rp/mygreeterv3/server/cmd/client hello --remote-addr <remote_addr>
```

#### Demoserver

To run the demoserver, you must use a different port than the server is already using, so you can send messages to the demoserver from the server.

To run the demoserver:

```bash
go run go.goms.io/aks/rp/mygreeterv3/server/cmd/demoserver start
```

To run the demoserver in a particular port:

```bash
go run go.goms.io/aks/rp/mygreeterv3/server/cmd/demoserver start --port <local_port>
```

#### Help

You can run help on every command in order to get more information on how to use them.

Examples:

```bash
go run go.goms.io/aks/rp/mygreeterv3/server/cmd/client help

go run go.goms.io/aks/rp/mygreeterv3/server/cmd/demoserver start -h
```


### Run service on AKS internal standalone

- Rename or delete your go.work file. Reason: aksbuilder doesn't work with go.work.
- [Create your standalone](https://dev.azure.com/msazure/CloudNativeCompute/_wiki/wikis/CloudNativeCompute.wiki/54089/Standalone-Environment-Usage)
- Make the targets in the order that they appear in the Makefile.

## Debugging and Common Failures


#### aksbuilder problems

- Clean up repo and aksbuilder cache.

```bash
git clean -xdf
sudo rm -rf ~/.aksbuilder 
```

## Monitoring

To view your service in Azure Data Explorer (ADX), follow [ADX dashboard creation/update instructions](server/monitoring/README.md).