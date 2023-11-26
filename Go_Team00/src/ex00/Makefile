generate:
	@cd messages && protoc --go_out=. --go-grpc_out=. proto/*.proto

run-server:
	@cd server && go run cmd/main.go

run-client:
	@cd client && go run cmd/main.go