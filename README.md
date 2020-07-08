# Planet Express

You are contracted to refactor Professor Farnsworth's legacy client/server code, and you decide to spice weasel (BAM!) things up with K8s, gRPC, and GraphQL. The professor wants to run a gateway service at the
Planet Express headquarters that can communicate with Planet Express ship via gRPC. Furthermore, he wants a nice React UI to show data about the ship, deliveries, crew members, etc that will communicate with the gateway service via GraphQL.

## Setup

When you are done shopping for lightspeed briefs you decide to setup the new project. You realize that some of it is already set up for you. The `./headquarters` directory contains the gateway/client code and the
`./ship` directory contains the server code.

1. Create a gitlab repo called `planet-express`
2. Setup Your Dev Environment
    - Install go (https://golang.org/doc/install)
    - Install protoc (https://grpc.io/docs/protoc-installation/)
    - Follow the getting started guide to install the protoc go plugin (https://grpc.io/docs/languages/go/quickstart/)
3. `go get google.golang.org/grpc`
4. `go get github.com/golang/protobuf`
5. Copy the files in this repo over to your new repo
6. Run `go mod init gitlab.com/<gitlab-username>/planet-express`
7. Your go mod should look something like the included `_go.mod` file (delete this file when done)
8. Fix the import statements in `./ship/main.go` and `./headquarters/main.go` (the lines commented like so `// pb "gitlab.com/<gitlab-username>/planet-express/ship/pkg/planetexpress"`)
9. Validate that you can build and genarate go code by running `make`
10. Validate you have the executables `./ship/ship` and `./headquarters/headquarters`
11. Run the server `./ship/ship` and the client `./headquarters/headquarters` in separate terminal windows and confirm the output.

## Generate Go Code

The protoc compiler with the go plugin will autogenerate client and server grpc stubs for you. Make sure this runs successfully with `make`

You should try changing the protobuf files in `./proto` and re-run `make`. The generated files are put in `./ship/pkg/planetexpress` and `./headquarters/pkg/planetexpress`

## Build

A makefile has been provided for you that will generate go code with the protoc compiler and build the executables. Run `make` and validate that you have executables at `./ship/ship` and `./headquarters/headquarters`

## Your Assignment

Time Estimate: 3 hours

The basic client and server code has been provided for you. Your job is to do the following:

1. Implement `crew.proto` and `delivery.proto` and the functions listed in `planet_express_service.proto`
    - You should experiment with different protobuf types for the crew and delivery resources. Each resource should demonstrate an understanding of protobuf.
2. Refactor `ship.proto` to use the `crew` and `delivery` resources.
3. Implement the rpc functions (defined in `planet_express_service.proto`) in both the client (`./headquarters`) and server (`./ship`), following the `GetShip` example. You do not need to add a database. Experiment with using data contained in the gRPC request.
4. Output data about the planet express ship, deliveries, and crew members to a json file called `planet_express.json`. Feel free to get creative with your data and protobuf resources.
5. Add a dockerfile to `./ship` so that we can run the server with docker.

### Above and Beyond

Time Estimate: 4 hours

1. Provide a basic helm template for the `./ship` server. This should contain a kubernetes deployment and service. We should be able to port-forward your service to call the `./ship` rpc functions.
2. Add additional protobuf resources like `ShipEngine`. You can get creative with this.
3. Add GraphQL support to the gateway (headquarters client). A good library for this is https://github.com/graph-gophers/graphql-go
4. Map a graphql resource and endpoint to a gRPC call such that you can use a GraphQL client to get data from the ship backend.

### Bite my shiny metal a$$ (extra, extra)

Time Estimate: 4 hours

1. Add a simple React SPA that displays data about the ship using GraphQL to communicate with the gateway (headquarters service). You can put this in the `./dashboard` directory.
2. Add some go unit tests to the ship server.

** These would really impress us here at Planet Express! **

## Deliverables
Make your repository public and share the repo url with Divvy's recruiter.

## Helpful Resources
- https://grpc.io/docs/languages/go/quickstart/
- https://grpc.io/docs/languages/go/basics/
- https://cloud.google.com/apis/design
- https://github.com/graph-gophers/graphql-go
- https://reactjs.org/docs/create-a-new-react-app.html

## Questions
If you find errors or otherwise need clarification on anything in the excercise, please reach out to Justin Sharp <justin.sharp@divvypay.com>.