protoc \
--proto_path=. \
--proto_path=third_party \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
--go-grpc_opt=require_unimplemented_servers=false \
--leo-grpc_out=. \
--leo-grpc_opt=paths=source_relative \
--leo-grpc_opt=require_unimplemented_servers=false \
api/*/*.proto