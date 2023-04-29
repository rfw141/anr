//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"github.com/rfw141/anr/gen/core"
	"github.com/rfw141/anr/pkg/app"
	"github.com/rfw141/anr/pkg/registry/etcd"
	"github.com/rfw141/anr/service/user/impl"
)

func initApp(*core.Config_Service, *core.Config_Registry) (*app.App, func(), error) {
	panic(wire.Build(
		etcd.ProviderSet,
		impl.ProviderSet,
		newApp,
	))
}
