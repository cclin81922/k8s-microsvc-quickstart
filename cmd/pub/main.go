package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "cloud.google.com/go/pubsub"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "golang.org/x/net/context"
    "google.golang.org/api/option"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

const (
    metrics_port = ":8080"
	svc_port = ":50051"
)

func metrics() {
    // expose prom /metrics
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(metrics_port, nil))
}

func svc(client interface{}) {
    // expose grpc service
	lis, err := net.Listen("tcp", svc_port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
    registerHelloworld(s, client)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func wait() {
    // wait for signal
    log.Print("Running...")
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    sig := <-sigs
    fmt.Printf(" %v\n", sig)
    log.Print("Exit")
}

func main() {
    // expose http api
    go metrics()

    // get gcloud pub/sub client object
    ctx := context.Background()
    client, err := pubsub.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEYJSON")))
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    } else {
        log.Printf("Client object %v", client)
    }

    // expose grpc api
    go svc(client)

    // wait for signal
    wait()
}
