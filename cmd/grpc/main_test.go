package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	grpcAuth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/gogo/googleapis/google/rpc"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestFlow(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := grpcAuth.NewAuthorizationClient(conn)

	request := &grpcAuth.CheckRequest{}

	ctx := context.TODO()

	response, err := client.Check(ctx, request)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)

	assert.Equal(t, int32(rpc.OK), response.GetStatus().GetCode())
}
