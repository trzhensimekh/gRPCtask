package middleware

import (
	"context"
	//"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func LoggerMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logrus.Info(req)
	logrus.Info(info)
	logrus.Println(ctxlogrus.Extract(ctx))
	return handler(ctx, req)
}
