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