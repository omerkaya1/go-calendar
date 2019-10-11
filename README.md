# go-calendar

Simple calendar of events. Receives new events passed to it and informs users of the upcoming
events (at least it should do so).

## Usage

```go-calendar [FLAGS]```

## Supported flags

### config (-c)
Path to the configuration file.

## TODO
First iteration (Calendar service, later on just CS):
##### Status: in progress
1) Define Event type; `DONE`
2) Create handlers for creating/updating/deleting events. `In Progress`
3) Keep events in an in-memory fashion (Use `map[string]Event`) `In Progress`

Second iteration (GRPC implementation for the CS):
##### Status: backlog
1) Lorem ipsum
2) Lorem ipsum

Third iteration (DB implementation for the the CS):
##### Status: somewhere near backlog
1) Lorem ipsum
2) Lorem ipsum

Fourth iteration (Message queue implementation for the the CS):
##### Status: deep backlog
1) Lorem ipsum
2) Lorem ipsum

Fifth iteration (Docker containerisation implementation for the the CS):
##### Status: nowhere near backlog
1) Lorem ipsum
2) Lorem ipsum

Sixth iteration (MtA vs. MsA implementation for the the CS):
##### Status: ...
1) Lorem ipsum
2) Lorem ipsum
