BASE_NAME := base
SERVER_NAME := server
CLIENT_NAME := client

generate:
	cd messages && protoc --go_out=. --go-grpc_out=. proto/*.proto

build-server:
	cd server && go build -o server ./cmd/

build-client:
	cd client && go build -o client ./cmd/

run-server:
	cd server && go run cmd/main.go

run-client:
	cd client && go run cmd/main.go

docker-base:
	docker build -t ${BASE_NAME} --file=./docker/Dockerfile.base ./

docker-server-build:
	docker build -t ${SERVER_NAME} --file=./server/build/Dockerfile ./

docker-client-build:
	docker build -t ${CLIENT_NAME} --file=./client/build/Dockerfile ./

docker-server-run:
	docker run -it --name server_app -p 4000:4000 --network grpc_network server

docker-client-run:
	docker run -d --rm --name client_app --network grpc_network client