run:
	go run main.go
run-server:
	go run service/grpc/server/main.go
proto-core: clen
	protoc --proto_path=proto proto/client.proto --go_out=service/grpc/ --go-grpc_out=service/grpc/
clen:
	rm -f service/grpc/oracle/*.pb.go
