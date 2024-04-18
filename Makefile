clean:
	rm pb/*
	rm swagger/*

gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb --openapiv2_out=:swagger

server1:
	go run server.go -port 50051

server2:
	go run server.go -port 50052

server:
	go run server.go -port 8080

rest:
	go run server.go -port 8081 -type rest -endpoint 0.0.0.0:8080

client:
	go run client.go -port 50051

test:
	go test -cover -race ./...

.PHONY: clean gen server client test cert 