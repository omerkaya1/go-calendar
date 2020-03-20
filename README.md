# go-calendar

Simple calendar of events. Receives new events passed to it and informs users of the upcoming
events (at least it should do so).

## Usage

```go-calendar [FLAGS]```

## Makefile targets and their description
- ```setup``` - Install all the build and lint dependencies
- ```mod``` - Runs mod
- ```test``` - Runs all unit tests
- ```integration-test``` - Runs integration tests
- ```cover``` - Runs all the tests and opens the coverage report
- ```fmt``` - Runs goimports on all go files
- ```lint``` - Runs all the linters
- ```gen``` - Triggers the protobuf code generation
- ```build-gcs``` - Builds the go-calendar project
- ```build-notification``` - Builds the notification project
- ```build-watcher``` - Builds the watcher project
- ```build-all``` - Builds all binaries of the project
- ```dockerbuild-gcs``` - Builds a docker image with the go-calendar project
- ```dockerpush-gcs``` - Publishes the docker image to the registry
- ```dockerbuild-notification``` - Builds a docker image with the notification project
- ```dockerpush-notification``` - Publishes the docker image to the registry
- ```dockerbuild-watcher``` - Builds a docker image with a project
- ```dockerpush-watcher``` - Publishes the docker image to the registry
- ```docker-compose-up``` - Runs docker-compose command to kick-start the infrastructure
- ```docker-compose-down``` - Runs docker-compose command to remove the turn down the infrastructure
- ```integration``` - Run integration tests
- ```clean``` - Remove temporary files
- ```help``` - Print this help message and exit

## Client API
The programme has support for both GRPC and REST API with similar command invocation.
```cgo
Run GRPC Web Service client

Usage:
  go-calendar grpc-client [command]

Examples:
  go-calendar grpc-client create -h

Available Commands:
  create      Create calendar event
  delete      Delete calendar event
  get         Get calendar event
  update      Update calendar event

Flags:
  -e, --end string     ending date and hour of the event
  -x, --expired        delete expired events for a user
  -h, --help           help for grpc-client
  -s, --host string    host address to connect to (default "127.0.0.1")
  -i, --id string      internal event id
  -l, --list           list all events belonging to a user
  -n, --note string    additional note related to the event
  -o, --owner string   owner of the event
  -p, --port string    port of the host (default "7070")
  -b, --start string   starting date and hour of the event
  -t, --title string   event name

Use "go-calendar grpc-client [command] --help" for more information about a command.
```

## Server part

```cgo
Run GRPC Server

Usage:
  go-calendar grpc-server [flags]

Examples:
# Initialise from configuration file
go-calendar grpc-server -c /path/to/config.json

# Initialise from parameters
go-calendar grpc-server --host=127.0.0.1 --port=7777 --log=2 --dbname=db_name --dbuser=username

Flags:
  -c, --config string       path to the configuration file
      --dbhost string       db host (default "127.0.0.1")
  -n, --dbname string       db name (default "test")
      --dbpassword string   db password
      --dbport string       db port (default "5432")
  -u, --dbuser string       db user
  -h, --help                help for grpc-server
  -s, --host string         host address (default "127.0.0.1")
  -l, --log int             changes log level (default 1)
  -p, --port string         host port (default "7070")
  -m, --sslmode string      ssl mode (default "disable")
```

## Description

Consult [docs](./docs) folder for further reference.
Project's progress is outlined in the [markdown file](./docs/go-calendar/README.md) file as well as the Prometheus
expressions useful for creating line graphs from the received metrics.

## TODO:
1. Refactoring:
    - DB methods;
    ~~- contextualise the app;~~
    - internal package;
2. Add UI:
    - Angular
    - Vue.js ???
3. Create a basic script that will ping the PostgreSQL DB instead of using a sleep
4. Add Kafka MQ implementation and modify the existing MQ interface;
5. Add Redis implementation;
6. Add Swagger implementation;
7. Add Meta linter to the project (golang-ci lint);
8. Error wrapping ???;
9. CI implementation.
