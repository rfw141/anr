package impl

import (
	"github.com/google/wire"
	user "github.com/rfw141/anr/gen/svc/user/v1"
)

var ProviderSet = wire.NewSet(
	user.NewUserClient,
	NewAdminService,
	NewServer,
)
