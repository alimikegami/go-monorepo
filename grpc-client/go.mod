module github.com/alimikegami/go-monorepo/grpc-client

go 1.23.3

require (
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250414145226-207652e42e2e // indirect
)

require github.com/alimikegami/go-monorepo/grpc-server v0.0.0

replace github.com/alimikegami/go-monorepo/grpc-server => ../grpc-server
