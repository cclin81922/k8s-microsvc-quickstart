package main

import (
    "os"
    "log"
    "cloud.google.com/go/pubsub"
    "golang.org/x/net/context"
    "google.golang.org/api/option"
    "google.golang.org/api/iterator"
)

func main() {
    ctx := context.Background()

    client, err := pubsub.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEYJSON")))
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    } else {
        log.Printf("Client object %v", client)

        //list_topics
        topics := client.Topics(ctx)
        log.Printf("Topics collection %v", topics)
        for {
            topic, err := topics.Next()
            if err == iterator.Done {
                break
            } else if err != nil {
                log.Fatalf("err %v", err)
                break
            } else {
                log.Printf("Topic %v", topic)
            }
        }
    }
}

/*
MESSAGE rpc error: code = NotFound desc = Requested project not found or user does not have access to it (project=). Make sure to specify the unique project identifier and not the Google Cloud Console display name.
ACTION TO DO you need export environment variable GCP_PROJECT

MESSAGE rpc error: code = PermissionDenied desc = User not authorized to perform this action.
ACTION TO DO you need assign an appropriate role i.e. "roles/pubsub.viewer" to the service account
*/
