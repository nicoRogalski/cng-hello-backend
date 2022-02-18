# cng-hello-backend

Example cloud native application in Go.

## How to run:
### Local
```
$ go run .\cmd\app.go
```
### docker-compose
```
$  docker build -f build/package/docker/Dockerfile -t cng-hello-backend .

$ cd test/docker/cng-hello-backend

$ docker-compose up
```

#

## Thougths on the project
- Is Go ready to be used in the cloud enterprise environment ?
- Can Go detach big ship backends like java ?

## Done
- Base go file structure (https://github.com/golang-standards/project-layout)
- Architecture (example: https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)
- Dockerfile
- Helm chart
- Configuration (https://github.com/spf13/viper)
- Logger for json logging (https://github.com/rs/zerolog)
- Rest-Api (https://github.com/gin-gonic/gin)
- Monitoring endpoint for prometheus (https://github.com/prometheus/client_golang)
- Added Grafana
- Tracing for jeager (https://github.com/open-telemetry/opentelemetry-go/tree/main/exporters/jaeger)
- Tracing of gin (https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/gin-gonic/gin/otelgin)

# 

## In progress:
- Health endpoints with readiness and live checks.
- Versioning

#

## To be done:
### Database
- Add postgress to docker-compose
- Setup Table with a migration tool
- Add Gorm and implement persitance 
- Add ressource only docker-compose for local development

### Security
- Add OIDC
- Rest security with middleware (JWT validation)

### Rest
- Add api as generated from openapi.yml 

### Tracing
- Add trace and span in logger

# 

## To be fixed:
- Viper does not automaticaly uses env vars

#

#### Watch out in the future:
- https://github.com/open-telemetry/opentelemetry-go
- GraphQL
- Messageing (Kafka, Nats, RabbitMQ ...)
