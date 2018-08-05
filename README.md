# For development on macOS

For local dev and test (NOT NEED docker nor k8s nor registry)

```
export PATH=$PATH:~/go/bin                          # or create .envrc file
export GITHUB_USER=<THIS GITHUB USER>

# install tools
brew install grpc                                   # command -v protoc
                                                    # command -v grpc_cli
                                                    # command -v grpc_python_plugin
                                                    # command -v grpc_php_plugin
go get -u github.com/golang/protobuf/protoc-gen-go  # command -v protoc-gen-go

# download source code
go get -d github.com/$GITHUB_USER/k8s-microsvc-quickstart

# compile source code
cd ~/go/src/github.com/$GITHUB_USER/k8s-microsvc-quickstart
make

# run to launch server
export GCP_PROJECT=<YOUR GCP PROJECT ID>
export GCP_KEYJSON=$PWD/key.json
./out/pub                                           # to interrupt, use ctrl + c or kill -2
                                                    # to terminate, use kill -15
                                                    # to test, use curl http://localhost:8080/metrics
                                                    # to test, use grpc_cli ls localhost:50051
                                                    # to test, use grpc_cli ls localhost:50051 helloworld.Greeter -l
                                                    # to test, use grpc_cli type localhost:50051 helloworld.HelloRequest
                                                    # to test, use grpc_cli call localhost:50051 SayHello "name: 'gRPC CLI'"
                                                    # to test, use go run cmd/tester/*.go
```

For local docker test (NOT NEED k8s nor registry)

```
# build docker image
make image

# run to launch server 
docker run --rm -v $PWD:/data -e GCP_PROJECT=$GCP_PROJECT -e GCP_KEYJSON=/data/key.json -p 8080:8080 -p 50051:50051 --name microsvc $GITHUB_USER/k8s-microsvc-quickstart:latest
```

For local k8s test (NOT NEED registry)

* NOTE use NodePort to test

```
kubectl create secret generic key-json --from-file=$PWD/key.json
kubectl create configmap gcp-project --from-literal=gcp-project-id=$GCP_PROJECT
helm install --name microsvc --set github.user=$GITHUB_USER k8s/LocalK8s/microsvc
                                                    # to test, use curl http://localhost:30080/metrics
                                                    # to test, use grpc_cli ls localhost:30051
                                                    # to test, use grpc_cli ls localhost:30051 helloworld.Greeter -l
                                                    # to test, use grpc_cli type localhost:30051 helloworld.HelloRequest
                                                    # to test, use grpc_cli call localhost:30051 SayHello "name: 'gRPC CLI'"
                                                    # NOTE by default, the range of valid node ports is 30000-32767

# cleanup
helm delete --purge microsvc
kubectl delete configmap gcp-project
kubectl delete secret key-json
```

For remote k8s test

* NOTE use ClusterIP to test

```
# NEED push image to remote registry
# NEED switch kube context

kubectl create secret generic key-json --from-file=$PWD/key.json
kubectl create configmap gcp-project --from-literal=gcp-project-id=$GCP_PROJECT
helm install --name microsvc --set github.user=$GITHUB_USER k8s/RemoteK8s/microsvc

# cleanup
helm delete --purge microsvc
kubectl delete configmap gcp-project
kubectl delete secret key-json
```

# To Append a service

1. create and implement pkg/pb/NEW-SERVICE/NEW-SERVICE.pb.go
2. create and implement pkg/server/NEW-SERVICE/server.go
3. create and implement cmd/pub/NEW-SERVICE.go
4. append this line "registerNEW-SERVICE(s, client)" in the function "svc" of the file "cmd/pub/main.go"
5. create and implement cmd/tester/NEW-SERVICE.go
6. append this line "testNEW-SERVICE()" in the function "main" of the file "cmd/tester/main.go"

# To Create pb for other programming languages

For python

```
make python                                         # to test, use pip install googleapis-common-protos grpcio && python python/greeter_client.py
```

For php

```
make php                                            # to test, use composer install && php php/greeter_client.php
                                                    # WIP
```

For js

```
make js                                             # to test, use npm install google-protobuf grpc && node js/greeter_client.js
```

For ruby

```
make ruby                                           # to test, use gem install ... && ruby ruby/greeter_client.rb
                                                    # WIP
```

# For release management with GCP

For docker image publishment

```
export TAG=v1.0.0

# download source code
git clone -b $TAG https://github.com/$GITHUB_USER/k8s-microsvc-quickstart

# build docker image
cd k8s-microsvc-quickstart
make image

# tag docker image
docker tag $GITHUB_USER/k8s-microsvc-quickstart:latest asia.gcr.io/$GCP_PROJECT/k8s-microsvc-quickstart:$TAG

# publish docker image
gcloud auth configure-docker
docker push asia.gcr.io/$GCP_PROJECT/k8s-microsvc-quickstart:$TAG       # to check, use gcloud container images list-tags asia.gcr.io/$GCP_PROJECT/k8s-microsvc-quickstart
                                                                        # to check, use gcloud container images describe asia.gcr.io/$GCP_PROJECT/k8s-microsvc-quickstart:$TAG
```

For rolling updates with k8s deployment

```
helm upgrade --name microsvc --set gcp.project=$GCP_PROJECT --set image.tag=$TAG k8s/ProdK8s/microsvc
```

# .envrc

```
export PATH=$PATH:~/go/bin
export GCP_PROJECT=xxx-project-xxxxxx
export GCP_KEYJSON=$PWD/key.json
export GITHUB_USER=xxx
```

Files, which support environment variables

* Dockerfile
* Makefile

Files, which use but don't support environment variables

* Helm Chart # workaround is using --set command-line argument
* grpc server/client code # workaround is using script to overwrite
  * hello_grpc/main.go
  * cmd/pub/helloworld.go
  * cmd/tester/helloworld.go

```
sed -i "s/cclin81922/$GITHUB_USER/g" hello_grpc/main.go cmd/pub/helloworld.go cmd/tester/helloworld.go pkg/server/helloworld/server.go
```
