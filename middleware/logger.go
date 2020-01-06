package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func LoggerMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error){
	fmt.Println("loging ...")
	return handler(ctx, req)
}
