compile:
	protoc --go_out=../build/pb/ \
			 --go-grpc_out=../build/pb/ \
			 --go_opt=paths=source_relative \
			 --go-grpc_opt=paths=source_relative \
			 ./*.proto ./service/noname.proto