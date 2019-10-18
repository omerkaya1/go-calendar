# go-calendar

Simple calendar of events. Receives new events passed to it and informs users of the upcoming
events (at least it should do so).

## Usage

```go-calendar [FLAGS]```

## Supported flags

### config (-c)
Path to the configuration file.

## TODO
### First iteration (Calendar service, later on just CS):
##### Status: in progress
1) Define Event type; `DONE`
2) Create handlers for creating/updating/deleting events. `DONE`
3) Keep events in an in-memory fashion (Use `map[string]Event`) `DONE`

### Second iteration (GRPC implementation for the CS):
##### Status: backlog
1) Add code generation target to the Makefile; `DONE`
2) Add client and service functionality (GRPC `DONE`, RESTful `DONE`)

### Third iteration (DB implementation for the the CS):
##### Status: somewhere near backlog
1) Integrate DB usage; `In Progress`
2) Make sure that the DB logic is independent from the higher level abstractions. `In Progress`

### Fourth iteration (Message queue implementation for the the CS):
##### Status: deep backlog
1) Lorem ipsum
2) Lorem ipsum

### Fifth iteration (Docker containerisation implementation for the the CS):
##### Status: nowhere near backlog
1) Lorem ipsum
2) Lorem ipsum

### Sixth iteration (MtA vs. MsA implementation for the the CS):
##### Status: ...
1) Lorem ipsum
2) Lorem ipsum
