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
##### Status: `DONE`
1) Define Event type; `DONE`
2) Create handlers for creating/updating/deleting events. `DONE`
3) Keep events in an in-memory fashion (Use `map[string]Event`) `DONE`

### Second iteration (GRPC implementation for the CS):
##### Status: `DONE`
1) Add code generation target to the Makefile; `DONE`
2) Add client and service functionality (GRPC `DONE`, RESTful `DONE`)

### Third iteration (DB implementation for the the CS):
##### Status: `DONE`
1) Integrate DB usage; `DONE`
2) Make sure that the DB logic is independent from the higher level abstractions. `DONE`

### Fourth iteration (Message queue implementation for the the CS):
##### Status: `DONE`
1) Added support for message queue; `DONE`
2) Queue message production and consumption processes were tested. `DONE`

### Fifth iteration (Docker containerisation implementation for the the CS):
##### Status: `DONE`
1) Added Dockerfile and docker-compose files for the project. `DONE`

### Sixth iteration (MtA vs. MsA implementation for the the CS):
##### Status: `DONE`
1) Restructure the code to implement microservice architecture

### Seventh iteration (Integration tests and Unit tests for the the CS):
##### Status: `DONE`
1) Add integration tests to the project;
2) Use GODOG to implement BDD testing approach.

### Eighth iteration (Prometheus monitoring implementation for the the CS):
##### Status: `DONE`
1) Add support for prometheus to the infrastructure;
2) Add metrics to the go-calendar, PostgreSQL and notification containers.

## Prometheus metrics
All metrics are available on localhost:9090 after `make docker-compose-up` is called.
### Go-calendar
##### Total number of RPCs completed on the server, regardless of success or failure.
```yaml
# TYPE grpc_server_handled_total counter
grpc_server_handled_total{grpc_code="OK",grpc_method="CreateEvent",grpc_service="GoCalendarServer",grpc_type="unary"}
grpc_server_handled_total{grpc_code="OK",grpc_method="GetEvent",grpc_service="GoCalendarServer",grpc_type="unary"}
grpc_server_handled_total{grpc_code="OK",grpc_method="UpdateEvent",grpc_service="GoCalendarServer",grpc_type="unary"}
grpc_server_handled_total{grpc_code="OK",grpc_method="DeleteEvent",grpc_service="GoCalendarServer",grpc_type="unary"}
grpc_server_handled_total{grpc_code="OK",grpc_method="GetUserEvents",grpc_service="GoCalendarServer",grpc_type="unary"}
```

### Postgres-Exporter
##### Number of sequential scans initiated on this table
```yaml
# TYPE pg_stat_user_tables_seq_scan counter
pg_stat_user_tables_seq_scan{datname="calendar_db",relname="events",schemaname="public",server="postgres:5432"}
```
##### Number of rows inserted
```yaml
# TYPE pg_stat_user_tables_n_tup_ins counter
pg_stat_user_tables_n_tup_ins{datname="calendar_db",relname="events",schemaname="public",server="postgres:5432"}
```
##### Estimated number of rows changed since last analyze
```yaml
# TYPE pg_stat_user_tables_n_mod_since_analyze gauge
pg_stat_user_tables_n_mod_since_analyze{datname="calendar_db",relname="events",schemaname="public",server="postgres:5432"}
```
##### Number of live rows fetched by index scans
```yaml
# TYPE pg_stat_user_tables_idx_tup_fetch counter
pg_stat_user_tables_idx_tup_fetch{datname="calendar_db",relname="events",schemaname="public",server="postgres:5432"}
```
##### Estimated number of live rows
```yaml
# TYPE pg_stat_user_tables_n_live_tup gauge
pg_stat_user_tables_n_live_tup{datname="calendar_db",relname="events",schemaname="public",server="postgres:5432"}
```
##### Number of index scans initiated on this table
```yaml
# TYPE pg_stat_user_tables_idx_scan counter
pg_stat_user_tables_idx_scan{datname="calendar_db",relname="events",schemaname="public",server="postgres:5432"}
```
##### Number of live rows fetched by sequential scans
```yaml
# TYPE pg_stat_user_tables_seq_tup_read counter
pg_stat_user_tables_seq_tup_read{datname="calendar_db",relname="events",schemaname="public",server="postgres:5432"}
```
##### Number of rows updated
```yaml
# TYPE pg_stat_user_tables_n_tup_upd counter
pg_stat_user_tables_n_tup_upd{datname="calendar_db",relname="events",schemaname="public",server="postgres:5432"}
```

### Notification
##### Total number of messages sent by the MQ service
```yaml
# TYPE rabbit_mq_sent_messages counter
rabbit_mq_sent_messages
```
