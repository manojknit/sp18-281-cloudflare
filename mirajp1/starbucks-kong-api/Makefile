
all: clean

clean: 
	find . -name 'goapi' -type f -exec rm -f {} \; 
	go clean

run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt goapi

install:
	go install goapi

build:
	go build goapi

start:
	./goapi 

test-ping:
	curl localhost:3000/ping

test-gumball:
	curl localhost:3000/gumball

docker-build: 
	docker build -t kong-api .
	docker images

docker-run:
	docker run -e KONG_SERVER=54.183.59.125 --name kong-api -td -p 3000:3000 kong-api
	docker ps

docker-network:
	docker network ls

docker-network-inspect:
	docker network inspect host

docker-shell:
	docker exec -it kong-api bash 

docker-clean:
	docker stop kong-api
	docker rm kong-api
	docker rmi kong-api

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"

