protoc \
		--proto_path=. \
		--proto_path=third_party \
		--go_out=. \
		--go_opt=paths=source_relative \
		*.proto