package reqlogger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func NewInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	l := log.With(zap.String("unary interceptor", "logger"))

	l.Info("added logger unary interceptor")

	logFn := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			l.Warn("no incoming metadata")
		}

		entry := l.With(
			zap.String("uri", info.FullMethod),
			zap.String("method", "unary"),
			zap.Any("req", req),
			zap.Any("metadata", md),
		)

		resp, err := handler(ctx, req)

		respMsg, ok := resp.(proto.Message)
		if ok {
			entry = entry.With(zap.Int("size", proto.Size(respMsg)))
		} else {
			log.Warn("unable to get response message")
		}

		entry.Info("request",
			zap.String("duration", time.Since(start).String()),
			zap.String("status", fmt.Sprintf("%d %s", int(status.Code(err)), status.Code(err))),
		)

		return resp, err
	}

	return logFn
}
