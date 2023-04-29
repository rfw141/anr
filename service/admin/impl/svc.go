package impl

import (
	"context"
	admin "github.com/rfw141/anr/gen/svc/admin/v1"
	user "github.com/rfw141/anr/gen/svc/user/v1"
	"log"
)

func NewAdminService(svcUser user.UserServiceClient) *AdminService {
	return &AdminService{
		svcUser: svcUser,
	}
}

type AdminService struct {
	admin.UnimplementedAdminServiceServer

	svcUser user.UserServiceClient
}

func (s *AdminService) CreateUser(ctx context.Context, req *admin.CreateUserReq) (*admin.CreateUserRsp, error) {
	var rsp admin.CreateUserRsp
	log.Printf("create user %+v", req)
	res, err := s.svcUser.CreateUser(ctx, &user.CreateUserReq{
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Printf("err:%v", err)
		return nil, err
	}
	log.Printf("res:%+v", res)
	return &rsp, nil
}
