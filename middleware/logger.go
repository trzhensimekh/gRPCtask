package middleware

import (
	"context"
	"fmt"

	//"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)


func LoggerMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("loging ...")
		logrus.Println(req)
		logrus.Println(ctxlogrus.Extract(ctx))
		return handler(ctx, req)
}