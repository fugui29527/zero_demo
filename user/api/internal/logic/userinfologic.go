package logic

import (
	"context"
	"zero_demo/user/rpc/pb"

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
	//userInfo := map[int]string{
	//	1: "lilei",
	//	2: "zhangshan",
	//}
	//username := "unkonw"
	//if name, ok := userInfo[req.Id]; ok {
	//	username = name
	//}
	//rpc调用
	in := &pb.GetUserInfoReq{
		Id: req.Id,
	}
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, in)
	if err != nil {
		return
	}
	resp = &types.Response{
		Name: user.Name,
	}
	return
}
