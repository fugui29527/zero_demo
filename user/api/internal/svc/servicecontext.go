package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero_demo/user/api/internal/config"
	"zero_demo/user/rpc/usercenter"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usercenter.NewUserCenter(zrpc.MustNewClient(c.UserRpc)),
	}
}
