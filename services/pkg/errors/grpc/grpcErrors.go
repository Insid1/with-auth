package grpc

import "errors"

var (
	ErrUnableToListenGrpcServer = errors.New("unable to listen GRPC server: ")
	ErrUnableToServeGrpcServer  = errors.New("unable to serve GRPC server: ")
)
