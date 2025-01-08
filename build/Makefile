registry="registry.cn-hangzhou.aliyuncs.com/daemon_public"
image_name="basic"
version=v1.0.0
server_name="basic"
project_name="basic"

build: build_proto
	@GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o $(server_name) ../cmd/$(project_name)
	@docker build . -t $(registry)/$(image_name):$(version)
	@rm -rf $(server_name)

build_proto:
	protoc -I ../api/proto \
           --go_out ../api/proto --go_opt paths=source_relative \
           --go-grpc_out=require_unimplemented_servers=false:../api/proto \
           --openapiv2_out ../api/docs \
           --grpc-gateway_out ../api/proto --grpc-gateway_opt paths=source_relative ../api/proto/*.proto

push:
	@docker push $(registry)/$(image_name):$(version)

release: build push
