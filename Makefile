BINARY_PATH=out/pub
PB_GO_PATH=pkg/pb/helloworld
PB_PYTHON_PATH=python
PB_PHP_PATH=php
PB_JS_PATH=js
PB_RUBY_PATH=ruby
IMAGE_TAG=latest

all: clean build

build:
	protoc -I pkg/pb/helloworld/ pkg/pb/helloworld/helloworld.proto --go_out=plugins=grpc:$(PB_GO_PATH)
	go build -o $(BINARY_PATH) github.com/$(GITHUB_USER)/k8s-microsvc-quickstart/cmd/pub

clean:
	rm -f $(PB_GO_PATH)/*.pd.go
	rm -f $(BINARY_PATH)

image:
	docker build --build-arg GITHUB_USER=$(GITHUB_USER) -t $(GITHUB_USER)/k8s-microsvc-quickstart:$(IMAGE_TAG) .

python:
	protoc -I pkg/pb/helloworld/ pkg/pb/helloworld/helloworld.proto --python_out=$(PB_PYTHON_PATH) --grpc_out=:$(PB_PYTHON_PATH) --plugin=protoc-gen-grpc=/usr/local/bin/grpc_python_plugin

php:
	protoc -I pkg/pb/helloworld/ pkg/pb/helloworld/helloworld.proto --php_out=$(PB_PHP_PATH) --grpc_out=:$(PB_PHP_PATH) --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin

js:
	# protoc -I pkg/pb/helloworld/ pkg/pb/helloworld/helloworld.proto --js_out=$(PB_JS_PATH) --grpc_out=:$(PB_JS_PATH) --plugin=protoc-gen-grpc=/usr/local/bin/grpc_node_plugin
	docker run --rm -v $(PWD)/pkg/pb:/pb -v $(PWD)/js:/js grpc/node:latest grpc_tools_node_protoc --js_out=import_style=commonjs,binary:/js --grpc_out=/js --plugin=protoc-gen-grpc=/usr/local/bin/grpc_tools_node_protoc_plugin --proto_path /pb/helloworld /pb/helloworld/helloworld.proto

ruby:
	protoc -I pkg/pb/helloworld/ pkg/pb/helloworld/helloworld.proto --ruby_out=$(PB_RUBY_PATH) --grpc_out=:$(PB_RUBY_PATH) --plugin=protoc-gen-grpc=/usr/local/bin/grpc_ruby_plugin

.PHONY: python php js ruby
