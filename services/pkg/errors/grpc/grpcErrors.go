package grpc

import "errors"

var (
	ErrUnableToListenGrpcServer = errors.New("Unable to listen GRPC server: ")
	ErrUnableToServeGrpcServer  = errors.New("Unable to serve GRPC server: ")
)
