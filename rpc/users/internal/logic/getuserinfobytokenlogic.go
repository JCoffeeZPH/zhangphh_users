package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"

	"zhangphh_users/rpc/users/internal/svc"
	"zhangphh_users/rpc/users/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByTokenLogic {
	return &GetUserInfoByTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByTokenLogic) GetUserInfoByToken(req *user.GetUserByTokenRequest) (*user.GetUserResponse, error) {
	token := req.GetToken()
	if len(token) == 0 {
		return nil, fmt.Errorf("token is null")
	}
	userId := parse(token)
	userInfo, err := l.svcCtx.Model.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	return &user.GetUserResponse{
		UserInfo: &user.User{
			UserId:   userId,
			Username: userInfo.Username,
		},
	}, nil
}

func (l *GetUserInfoByTokenLogic) getJwtToken(secretKey string, seconds, userId int64) (string, error) {
	createTime := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["expire_time"] = createTime + seconds // 过期时间
	claims["create_time"] = createTime           // token颁发时间
	claims["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func parse(token string) int64 {
	now := time.Now().Unix()
	t, err := jwt.Parse(token, nil)
	if t == nil {
		logx.Errorf("parseJwtToken, invalid token: %s, err: %+v", token, err)
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		logx.Infof("userId: %+v", claims["userId"])
		if claims["expire_time"].(float64) < float64(now) {
			return 1
		}
		// todo 续期
	}
	return -1
}
