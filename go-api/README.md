# Sample GO API

Main features:

## Requirements:

- Go 1.23
- Docker
- Docker Compose

## To run:

```SHELL
  make compose
```

OR

```SHELL
  make compose-rebuild
```

## Endpoints:

- GET /metrics

```SHELL
  curl http://localhost:8080/metrics
```

- GET /payments

```SHELL
  curl http://localhost:8080/payments -H "Status: success"
  curl http://localhost:8080/payments -H "Status: success"
``` 

DONE:

- Prometheus and Grafana Integrated
- Graceful Shutdown Implementation

WIP:
- API integrate

--- 

TODO:

- Mongo Integrate
- HealthCheck
- Postgrest Integrate
- Stress Test ( Compare with others API's)
- Tests with TestContainers
- Message Integration (Rabbit/Kafka)
- Cache (with TTL and Redis)
- Swagger