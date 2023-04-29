package impl

import (
	"github.com/rfw141/anr/gen/core"
	"github.com/rfw141/anr/gen/svc/user/v1"
	"github.com/rfw141/anr/pkg/server/grpc"
)

func NewServer(c *core.Config_Service, svc *UserSvc) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Middleware(),
	}
	if c.Network != "" {
		opts = append(opts, grpc.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, grpc.Address(c.Addr))
	}
	if c.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServiceServer(srv, svc)
	return srv
}
