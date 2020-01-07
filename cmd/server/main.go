package main

import (
	mw "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/trzhensimekh/cursesGo/gRPCtask/internal/data"
	"github.com/trzhensimekh/cursesGo/gRPCtask/internal/server"
	"github.com/trzhensimekh/cursesGo/gRPCtask/middleware"
	pb "github.com/trzhensimekh/cursesGo/gRPCtask/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":8081"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(mw.WithUnaryServerChain(middleware.LoggerMiddleware))
	srv:=server.NewServer(data.Open())
	pb.RegisterUserServiceServer(s,srv )
	defer srv.CloseDb()
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}