# cng-hello-backend

Example cloud native application in Go.

## 1. How to run:

### Local
```
go run .\cmd\app.go
```

### Local with database
```
cd test/docker/cng-hello-backend-ressources
docker-compose up
cd ../../../
go run .\cmd\app.go
```

### docker-compose
```
docker build -f build/package/docker/Dockerfile -t cng-hello-backend .

cd test/docker/cng-hello-backend-standalone

docker-compose up
```

#

## 2. Thougths on the project
- Is Go ready to be used in the cloud enterprise environment ?
- Can Go detach big ship backends like java ?

### 20.02.2022
- Some general packages for tracing, logging, and auth with a clean api needs to be implemented in order to not always do everything from scratch


## 3. Done
### General
- Base go file structure (https://github.com/golang-standards/project-layout)
- Architecture (example: https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)
- Rest-Api (https://github.com/gin-gonic/gin)
- Logger for json logging (https://github.com/rs/zerolog)

### Database
- Added usage of gorm for postgresql (https://github.com/go-gorm/gorm)

### Build/Deployment
- Dockerfile
- Helm chart
- Configuration (https://github.com/spf13/viper)

### Auth
- JWT validation (https://github.com/golang-jwt/jwt)
- Cert parsing from oidc cert endpoint
- Apply groups from jwt in the context

### Metrics
- Monitoring endpoint for prometheus (https://github.com/prometheus/client_golang)
- Grafana
- Tracing for jeager (https://github.com/open-telemetry/opentelemetry-go/tree/main/exporters/jaeger)
- Tracing of gin (https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/gin-gonic/gin/otelgin)
- Trace and spanId in logs
- Implement readiness and live checks

### Pkg
- Added pgk for public usage of default implementations for gin/jwt/oidc (Should be later generalized and outsourced)

# 

## 4. In progress:

### Database
- Setup Table with a migration tool

### Monitoring
- Setup grafana that the dashboard is predefined

#

## 5. To be done:

### Rest
- Add api as generated from openapi.yml 
- Routing should be done in specific handlers

### Testing
- Apply testing for endpoints and services

# 

## 6. Known issues:
Nothing so far

#

## 7. Future concepts
- https://github.com/open-telemetry/opentelemetry-go
- GraphQL
- Messaging (Kafka, Nats, RabbitMQ ...)
- Outsource pkg in different project as lib/module/package
