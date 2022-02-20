# cng-hello-backend

Example cloud native application in Go.

## How to run:
### Local
```
go run .\cmd\app.go
```
### docker-compose
```
docker build -f build/package/docker/Dockerfile -t cng-hello-backend .

cd test/docker/cng-hello-backend

docker-compose up
```

#

## Thougths on the project
- Is Go ready to be used in the cloud enterprise environment ?
- Can Go detach big ship backends like java ?

### 20.02.2022
- Some general packages for tracing, logging, and auth with a clean api needs to be implemented in order to not always do everything from scratch


## Done
- Base go file structure (https://github.com/golang-standards/project-layout)
- Architecture (example: https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)
- Dockerfile
- Helm chart
- Configuration (https://github.com/spf13/viper)
- Logger for json logging (https://github.com/rs/zerolog)
- Rest-Api (https://github.com/gin-gonic/gin)
- Monitoring endpoint for prometheus (https://github.com/prometheus/client_golang)
- Grafana
- Tracing for jeager (https://github.com/open-telemetry/opentelemetry-go/tree/main/exporters/jaeger)
- Tracing of gin (https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/gin-gonic/gin/otelgin)
- Trace and spanId in logs
- JWT validation (https://github.com/golang-jwt/jwt)
- Cert parsing from oidc cert endpoint
- Apply groups from jwt in the context
- Implement readiness and live checks
- Added pgk for public usage of default implementations for gin

# 

## In progress:

### Monitoring
- Setup grafana that the dashboard is predefined

### Pkg
- add more default implementations that can be reused in pgk
- Add gin jwt middleware with options 

#

## To be done:

### Database
- Setup Table with a migration tool
- Add Gorm and implement persitance 

### Rest
- Add api as generated from openapi.yml 

# 

## Known issues:

#

#### Watch out in the future:
- https://github.com/open-telemetry/opentelemetry-go
- GraphQL
- Messageing (Kafka, Nats, RabbitMQ ...)
- outsource pkg in different project as lib/module/package
