package logic

import (
	"context"

	"zero_demo/user/rpc/internal/svc"
	"zero_demo/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	userInfo := map[int64]string{
		1: "lisi",
		2: "zhangsan",
	}
	username := "unknown"
	if name, ok := userInfo[in.Id]; ok {
		username = name
	}
	return &pb.GetUserInfoResp{
		Id:   in.Id,
		Name: username,
	}, nil
}
