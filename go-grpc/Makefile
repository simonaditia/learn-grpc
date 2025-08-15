# GO_WORKSPACE := $(GOPATH)/src
# protoc --proto_path=protos protos/*.proto --go_out=$(GO_WORKSPACE) --go-grpc_out=$(GO_WORKSPACE)

protoc:	
	protoc --proto_path=protos protos/*.proto --go_out=. --go-grpc_out=.
	@echo "Protoc compile selesai"