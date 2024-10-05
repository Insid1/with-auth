package pkg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Insid1/go-auth-user/pkg/grpc/auth_v1"
	"github.com/Insid1/go-auth-user/pkg/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthUnaryInterceptor интерцептор для проверки авторизации и аунтефикации
func AuthUnaryInterceptor(client auth_v1.AuthV1Client, methodsName []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		// check is method in methodsName
		if !utils.IsInList(info.FullMethod, methodsName) {
			return handler(ctx, req)
		}

		// Извлечение headers из запроса данных
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return "", status.Errorf(codes.Unauthenticated, "error retrieving headers")
		}

		// Проверка, что сервис имеет доступ к данным
		err = checkAuthService(md)
		if err == nil {
			return handler(ctx, req)
		}

		accessToken, err := retrieveAccessToken(md)
		if err != nil {
			return "", status.Errorf(codes.Unauthenticated, "error retrieving token")
		}

		_, err = client.Check(ctx, &auth_v1.CheckReq{
			AccessToken: accessToken,
		})

		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "error validating token")
		} else {
			return handler(ctx, req)
		}
	}
}

func retrieveAccessToken(md metadata.MD) (string, error) {
	// Извлечение авторизационных данных
	authDataList, ok := md["authorization"]
	if !ok {
		return "", fmt.Errorf("unable to retrieve auth data")
	}

	// Извлечение токена
	var accessToken string
	prefix := "Bearer "
	for _, authData := range authDataList {
		accessToken, ok = strings.CutPrefix(authData, prefix)
		if !ok {
			return "", fmt.Errorf("unable to retrieve auth data")
		}
	}

	return accessToken, nil
}

func checkAuthService(md metadata.MD) error {
	// Извлечение авторизационных данных
	serviceDataList, ok := md["service-name"]
	if !ok {
		return fmt.Errorf("unable to retrieve service info")
	}

	// Извлечение токена
	for _, serviceName := range serviceDataList {
		if serviceName == "auth" {
			return nil
		}
	}

	return fmt.Errorf("not auth service")
}
