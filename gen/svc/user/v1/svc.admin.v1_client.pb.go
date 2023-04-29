package v1

import (
	"context"
	"github.com/rfw141/anr/internal"
	"github.com/rfw141/anr/pkg/server/grpc"
)

func NewUserClient(discovery internal.Discovery) UserServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user"),
		grpc.WithDiscovery(discovery),
	)
	if err != nil {
		panic(err)
	}
	c := NewUserServiceClient(conn)
	return c
}
