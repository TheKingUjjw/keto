package main

import (
	"context"
	"fmt"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4466", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err.Error())
	}

	client := rts.NewCheckServiceClient(conn)

	res, err := client.Check(context.Background(), &rts.CheckRequest{
		Namespace: "messages",
		Object:    "02y_15_4w350m3",
		Relation:  "decypher",
		Subject:   rts.NewSubjectID("john"),
	})
	if err != nil {
		panic(err.Error())
	}

	if res.Allowed {
		fmt.Println("Allowed")
		return
	}
	fmt.Println("Denied")
}
