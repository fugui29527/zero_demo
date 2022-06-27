package logic

import (
	"context"

	"zero_demo/user/api/internal/svc"
	"zero_demo/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	userInfo := map[int]string{
		1: "lilei",
		2: "zhangshan",
	}
	username := "unkonw"
	if name, ok := userInfo[req.Id]; ok {
		username = name
	}
	resp = &types.Response{
		Name: username,
	}
	return
}
