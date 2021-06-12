protoc \
-I . \
-I ../../../grpc-gateway/ \
-I ${GOPATH}/src/github.com/googleapis/ \
--go_out=paths=source_relative:./ \
--go-grpc_out=paths=source_relative:./ \
--grpc-gateway_out=logtostderr=true,allow_patch_feature=false,paths=source_relative,allow_delete_body=true:./ \
./test.proto