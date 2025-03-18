protoc \
		--proto_path=. \
		--proto_path=../../proto/ \
		--proto_path=../../third_party \
		--go_out=. \
		--go_opt=paths=source_relative \
		--validate_out=. \
		--validate_opt=paths=source_relative,lang=go \
		--go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
		--go-leo_out=. \
		--go-leo_opt=paths=source_relative \
		*/*/*.proto