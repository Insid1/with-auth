package pkg

import (
	"context"

	"github.com/Insid1/go-auth-user/pkg/grpc/auth_v1"
	clientInterceptors "github.com/Insid1/go-auth-user/pkg/interceptors/client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCInitializedAuthClient struct {
	Client     auth_v1.AuthV1Client
	Connection *grpc.ClientConn
}

type GRPCAuthClientConfig struct {
	ServerAddress     string
	ClientServiceName string
}

func InitGRPCAuthClient(ctx context.Context, cfg *GRPCAuthClientConfig) (*GRPCInitializedAuthClient, error) {
	connection, err := grpc.NewClient(
		cfg.ServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			clientInterceptors.UnaryServiceInfoInterceptor(cfg.ClientServiceName),
		),
	)
	if err != nil {
		return nil, err
	}

	return &GRPCInitializedAuthClient{
		Client:     auth_v1.NewAuthV1Client(connection),
		Connection: connection,
	}, nil
}
