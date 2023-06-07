
dfp := build/package/docker/Dockerfile
rcfp := deployment/docker/cng-hello-backend-resources/docker-compose.yml
scfp := deployment/docker/cng-hello-backend-standalone/docker-compose.yml

app:
	@go build -o bin/app ./

run: app
	@./bin/app -local

compose-resources-up: 
	docker compose -f $(rcfp) up --detach

compose-resources-down:
	docker compose -f $(rcfp) down

image:
	docker build -f $(dfp) -t cng-hello-backend --build-arg version=$$(cat VERSION) .

compose-standalone-up: image
	docker compose -f $(scfp) up --detach

compose-standalone-down:
	docker compose -f $(scfp) down
	
# pprof
pprof-heap: 
	curl http://localhost:8080/debug/pprof/heap > heap.out

pprof-allocs: 
	curl http://localhost:8080/debug/pprof/allocs > allocs.out

pprof-heap-web: pprof-heap
	go tool pprof -http=:8082 heap.out

pprof-allocs-web: pprof-allocs
	go tool pprof -http=:8082 allocs.out
