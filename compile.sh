
cd ./proto/ && \
protoc -I . --go_out=. \
    --grpc-gateway_out=. \
    --go-grpc_out=. \
    --go_opt=paths=source_relative \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --go-grpc_opt=paths=source_relative \
    *.proto 
