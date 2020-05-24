
protogen:
	protoc --proto_path=api/protobuf --go_out=plugins=grpc:. api/protobuf/*.proto

run-grpc-server:
	go run cmd/grpc-server/main.go $(profile)

run-grpc-client:
	go run cmd/grpc-client/main.go $(profile)