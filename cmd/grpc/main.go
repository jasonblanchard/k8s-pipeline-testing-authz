package main

import (
	"context"
	"fmt"
	"net"

	grpcAuth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/gogo/googleapis/google/rpc"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
)

// AuthzServer handles envoy ext_authz grpc requests
type AuthzServer struct{}

// Check do the request check
func (s *AuthzServer) Check(ctx context.Context, request *grpcAuth.CheckRequest) (*grpcAuth.CheckResponse, error) {
	fmt.Println(request.GetAttributes().GetSource())
	fmt.Println(request.GetAttributes().GetDestination())
	fmt.Println(request.GetAttributes().GetRequest().GetHttp())

	return &grpcAuth.CheckResponse{
		Status: &status.Status{
			Code: int32(rpc.OK),
		},
		HttpResponse: &grpcAuth.CheckResponse_OkResponse{
			OkResponse: &grpcAuth.OkHttpResponse{
				// Headers: headers,
			},
		},
	}, nil
}

func main() {
	s := &AuthzServer{}

	port := "8080"

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	fmt.Println(fmt.Sprintf("Listening on port %s", port))

	grpcServer := grpc.NewServer()
	grpcAuth.RegisterAuthorizationServer(grpcServer, s)
	grpcServer.Serve(lis)
}
