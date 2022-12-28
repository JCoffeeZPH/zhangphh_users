package logic

import (
	"context"

	"zhangphh_users/rpc/users/internal/svc"
	"zhangphh_users/rpc/users/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *user.GetUserRequest) (*user.GetUserResponse, error) {
	u, e := l.svcCtx.Model.FindOne(l.ctx, req.GetUserId())
	if e != nil {
		panic(e)
	}

	return &user.GetUserResponse{
		UserInfo: &user.User{
			UserId:   u.Id,
			Username: u.Username,
		},
	}, nil
}
