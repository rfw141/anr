package etcd

import (
	"github.com/google/wire"
	"github.com/rfw141/anr/gen/core"
	"github.com/rfw141/anr/internal"
	etcd "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewEtcd, NewDiscovery, NewRegistrar)

func NewEtcd(config *core.Config_Registry) *etcd.Client {
	client, err := etcd.New(etcd.Config{
		Endpoints: config.Etcd.Endpoints,
	})
	if err != nil {
		panic(err)
	}
	return client
}

func NewDiscovery(cli *etcd.Client, config *core.Config_Registry) internal.Discovery {
	return New(cli)
}

func NewRegistrar(cli *etcd.Client, config *core.Config_Registry) internal.Registrar {
	return New(cli)
}
