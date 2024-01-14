PB_DIR := pb
PROTO_DIR := proto
PYTHON_DIR := .

app:
	grpc_tools_node_protoc \
		--js_out=import_style=commonjs,binary:${PB_DIR} \
		--grpc_out=grpc_js:${PB_DIR} \
		${PROTO_DIR}/app.proto
	protoc \
		--go_out=${PB_DIR} \
		--go_opt=paths=source_relative \
		--go-grpc_out=${PB_DIR} \
		--go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/app.proto

quant:
	protoc \
		--go_out=${PB_DIR} \
		--go_opt=paths=source_relative \
		--go-grpc_out=${PB_DIR} \
		--go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/quant.proto
	python3 -m grpc_tools.protoc \
		--proto_path=${PROTO_DIR} \
		--python_out=${PYTHON_DIR} \
		--grpc_python_out=${PYTHON_DIR} \
		${PROTO_DIR}/quant.proto

gen: app quant
