package common

import "google.golang.org/grpc"

type GRPCClient[C interface{}] struct {
	Connection *grpc.ClientConn
	Client     C
}
