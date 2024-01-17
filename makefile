GO_DIR := go
NODE_DIR := node
PYTHON_DIR := python
PROTO_DIR := proto

app:
	grpc_tools_node_protoc \
		--proto_path=${PROTO_DIR} \
		--js_out=${NODE_DIR} \
		--grpc_out=${NODE_DIR} \
		${PROTO_DIR}/app.proto
	protoc \
		--proto_path=${PROTO_DIR} \
		--go_out=${GO_DIR} \
		--go_opt=paths=source_relative \
		--go-grpc_out=${GO_DIR} \
		--go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/app.proto

quant:
	protoc \
		--proto_path=${PROTO_DIR} \
		--go_out=${GO_DIR} \
		--go_opt=paths=source_relative \
		--go-grpc_out=${GO_DIR} \
		--go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/quant.proto
	python3 -m grpc_tools.protoc \
		--proto_path=${PROTO_DIR} \
		--python_out=${PYTHON_DIR} \
		--grpc_python_out=${PYTHON_DIR} \
		${PROTO_DIR}/quant.proto

gen: app quant
