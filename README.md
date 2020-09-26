# NoName - GRPC SDK #

protoc \
  --proto_path=./import/ \
  --go_out=./export/proto \
  --go-grpc_out=./export/proto \
  ./import/proto/*.proto