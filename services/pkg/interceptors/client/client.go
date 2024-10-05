package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// интерцептор для добавления информации о сервисе в метаданные при запросе
func UnaryServiceInfoInterceptor(serviceName string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Добавляем информацию о сервисе в метаданные
		md := metadata.Pairs("service-name", serviceName)
		ctx = metadata.NewOutgoingContext(ctx, md)

		metadata.FromIncomingContext(ctx)

		// Вызываем следующий обработчик в цепочке
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
