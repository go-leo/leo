protoc \
--proto_path=. \
--proto_path=../third_party \
--proto_path=../../ \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
--go-grpc_opt=require_unimplemented_servers=false \
--leo_out=. \
--leo_opt=paths=source_relative \
api/*/*.proto \
configs/*.proto
