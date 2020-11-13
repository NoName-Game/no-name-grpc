# NoName - GRPC SDK #

### Protobuf standard
```
protoc \
  --proto_path=./import/ \
  --go_out=./export/proto \
  --go-grpc_out=./export/proto \
  ./import/proto/*.proto
```
  
### Protobuf gogo
```
  protoc \
  -I=./import/ \
  -I=/go/src/github.com/gogo/protobuf/protobuf \
  --gogofast_out=Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,plugins=grpc:./export/pb \
  ./import/proto/*.proto ./import/noname.proto"
```