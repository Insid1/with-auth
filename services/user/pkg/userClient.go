package pkg

import (
	"context"

	"github.com/Insid1/go-auth-user/pkg/grpc/user_v1"

	clientInterceptors "github.com/Insid1/go-auth-user/pkg/interceptors/client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCInitializedUserClient struct {
	Client     user_v1.UserV1Client
	Connection *grpc.ClientConn
}

type GRPCUserClientConfig struct {
	ServerAddress     string
	ClientServiceName string
}

func InitGRPCUserClient(ctx context.Context, cfg *GRPCUserClientConfig) (*GRPCInitializedUserClient, error) {
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

	return &GRPCInitializedUserClient{
		Client:     user_v1.NewUserV1Client(connection),
		Connection: connection,
	}, nil
}
