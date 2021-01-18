# migration-test

## About

Playground script to test Mongo migrations using `migrate` package.

## Usage of this project

### Spin up environment

```
docker-compose up -d
```

### Example Script

```
# migrate up all
go run cmd/mongo/main.go -up

# migrate down all
go run cmd/mongo/main.go -down

# migrate up first 2 steps
go run cmd/mongo/main.go -up -steps 2

# migrate down most recent 2 steps
go run cmd/mongo/main.go -down -steps 2
```

## Other usage of package

### CLI
```
migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2
```

### Docker

```
docker run -v {{ migration dir }}:/migrations --network host migrate/migrate \
    -path=/migrations/ -database postgres://localhost:5432/database up 2
```
