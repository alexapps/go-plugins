// Package grpc provides a grpc transport
// Deprecated: use `github.com/alexapps/go-micro/transport/grpc` instead
package grpc

import (
	"github.com/alexapps/go-micro/transport"
	"github.com/alexapps/go-micro/transport/grpc"
)

// Deprecated: use `github.com/alexapps/go-micro/transport/grpc` instead
func NewTransport(opts ...transport.Option) transport.Transport {
	return grpc.NewTransport(opts...)
}
