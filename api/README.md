# Backend API

## Stack

- Go 1.26
- Gin
- ClickHouse
- Prometheus
- Grafana
- JWT Authentication
- Generic Pipeline Library

## Project Structure

```
cmd/
internal/
pkg/
docker/
```

## Running

```bash
cp .env.example .env

go mod tidy

make run
```

## Docker

```bash
make docker
```

## Endpoints

GET /health

GET /metrics

## Upcoming

- JWT Authentication

- Event Pipeline

- ClickHouse Analytics

- Grafana Dashboard