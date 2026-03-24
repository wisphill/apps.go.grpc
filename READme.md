### Prerequisites

```
# protobuf
brew install protobuf
# go builder Docker, using dev-toolchain/go-builder:1.0
make proto
```

### Architecture

```
order-service/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── domain/
│   │   └── order.go
│   ├── service/
│   │   └── order_service.go
│   ├── repository/
│   │   └── order_repo.go
│   ├── interceptor/
│   │   ├── auth.go
│   │   └── logging.go
│   └── transport/
│       └── grpc/
│           └── handler.go
├── proto/
│   └── order.proto
├── pkg/
│   └── logger/
├── configs/
├── Dockerfile
├── docker-compose.yml
└── go.mod

```

### Useful commands

```
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext \
  -H "authorization: Bearer test-token" \
  -d '{"userId":"1","amount":"200"}' \
  localhost:50051 \
  order.OrderService/CreateOrder

```
