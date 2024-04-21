package interceptors

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"mail/internal/pkg/logger"
	"os"
	"runtime/debug"
)

// PanicRecoveryInterceptor intercepts panics, recovers, logs info, and sets up logging and requestID.
func PanicRecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if errPanic := recover(); errPanic != nil {
			log.Println(
				"Panic, ",
				"Method: ", info.FullMethod,
				//"Error: ", errPanic.(string),
				"Message: ", string(debug.Stack()),
			)
		}
	}()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata error")
	}

	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "log.txt")
	}
	defer f.Close()

	ctx2 := context.WithValue(ctx, "logger", logger.InitializationBdLog(f))
	ctx3 := context.WithValue(ctx2, "requestID", md.Get("requestID"))

	return handler(ctx3, req)
}
