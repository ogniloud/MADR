package protobuf

//go:generate protoc --go_out=../internal/wordmaster --go-grpc_out=../internal/wordmaster --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./requests.proto ./common.proto ./responses.proto ./spanish.proto
