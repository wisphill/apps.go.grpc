COMPOSE=docker

PROTO_DIR=proto
OUT_DIR=.

PROJECT_DIR=$(CURDIR)

.PHONY: proto

proto:
	$(COMPOSE) run --rm \
	-v $(PROJECT_DIR):/workspace \
	-w /workspace \
	dev-toolchain/go-builder:1.0 sh -c "\
	find proto -name '*.proto' | xargs protoc \
	--proto_path=proto \
	--go_out=. \
	--go-grpc_out=."
