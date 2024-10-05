package utils

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Для обработки паники внутри запросов к серверу
func GetUnaryPanicInterceptor(lgr *zap.SugaredLogger) grpc.UnaryServerInterceptor {

	// Логирование паники
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			// Логируем информацию о панике с уровнем Error
			lgr.Errorf("Recovered from panic. panic: %s", p)

			// При паники возвращает internal error пользователю
			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	return recovery.UnaryServerInterceptor(recoveryOpts...)
}

// интерцептор для добавления информации о сервисе в метаданные
func GetUnaryServiceInfoInterceptor(serviceName string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Добавляем информацию о сервисе в метаданные
		md := metadata.Pairs("service-name", serviceName)
		ctx = metadata.NewOutgoingContext(ctx, md)
		println(md["service-name"])

		metadata.FromIncomingContext(ctx)

		// Вызываем следующий обработчик в цепочке
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
