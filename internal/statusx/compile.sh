protoc \
		--proto_path=. \
		--proto_path=../../third_party \
		--proto_path=../../../ \
		--go_out=. \
		--go_opt=paths=source_relative \
		error.proto