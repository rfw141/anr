package impl

import (
	v1 "github.com/rfw141/anr/gen/svc/user/v1"
)

func NewUserSvc() *UserSvc {
	return &UserSvc{

	}
}

type UserSvc struct {
	v1.UnimplementedUserServiceServer
}
