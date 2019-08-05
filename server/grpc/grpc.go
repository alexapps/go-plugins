// Package grpc provides a grpc server
// Deprecated: use `github.com/alexapps/go-micro/server/grpc` instead
package grpc

import (
	"github.com/alexapps/go-micro/server"
	"github.com/alexapps/go-micro/server/grpc"
)

var (
	// DefaultMaxMsgSize define maximum message size that server can send
	// or receive.  Default value is 4MB.
	// Deprecated: use `github.com/alexapps/go-micro/server/grpc` instead
	DefaultMaxMsgSize = grpc.DefaultMaxMsgSize
)

// Deprecated: use `github.com/alexapps/go-micro/server/grpc` instead
func NewServer(opts ...server.Option) server.Server {
	return grpc.NewServer(opts...)
}
