module github.com/dawit_hopes/grpc_micro_service/client

go 1.23.0

toolchain go1.23.10

require (
	github.com/caarlos0/env/v10 v10.0.0
	github.com/dawit_hopes/grpc_micro_service/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.73.0
)

require (
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk v1.36.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/dawit_hopes/grpc_micro_service/proto => ../proto
