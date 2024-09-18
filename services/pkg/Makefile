LOCAL_BIN:=$(CURDIR)/bin

user_grpc_result_path="grpc/user_v1"
auth_grpc_result_path="grpc/auth_v1"
team_grpc_result_path="grpc/team"

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

all: generate-user-api generate-auth-api

generate-user-api:
	mkdir -p $(user_grpc_result_path);
	protoc --proto_path api/user \
	-I api \
	--go_out=$(user_grpc_result_path) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=$(user_grpc_result_path) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user/v1.proto

generate-auth-api:
	mkdir -p $(auth_grpc_result_path);
	protoc --proto_path api/auth \
	-I api \
	--go_out=$(auth_grpc_result_path) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=$(auth_grpc_result_path) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/auth/v1.proto
