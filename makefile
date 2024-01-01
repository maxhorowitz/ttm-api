PB_DIR := pb
PROTO_DIR := proto

gen:
	rm -rf ${PB_DIR} && mkdir ${PB_DIR}
	grpc_tools_node_protoc \
		--js_out=import_style=commonjs,binary:${PB_DIR} \
		--grpc_out=grpc_js:${PB_DIR} \
		${PROTO_DIR}/service.proto
	protoc \
		--go_out=${PB_DIR} \
		--go_opt=paths=source_relative \
		--go-grpc_out=${PB_DIR} \
		--go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/service.proto
