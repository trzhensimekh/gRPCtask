package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/trzhensimekh/cursesGo/gRPCtask/pb"
	"google.golang.org/grpc"
	"net/http"
)


func run(ep string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, ep, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	endpoint:= flag.String("endpoint", "localhost:8081", "address")
	flag.Parse()

	defer glog.Flush()

	if err := run(*endpoint); err != nil {
		glog.Fatal(err)
	}
}
