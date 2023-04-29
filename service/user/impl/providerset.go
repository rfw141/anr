package impl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserSvc,
	NewServer,
)
