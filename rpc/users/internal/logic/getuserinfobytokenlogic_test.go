package logic

import (
	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
	"time"
)

func TestGetUserInfoByTokenLogic_getJwtToken(t *testing.T) {
	createTime := int32(time.Now().Unix())
	t.Log(createTime)
	claims := make(jwt.MapClaims)
	claims["expire_time"] = createTime + 300 // 过期时间
	claims["create_time"] = createTime       // token颁发时间
	claims["userId"] = 1
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tt, _ := token.SignedString([]byte("abcdefhijklmn"))
	t.Log(tt)
	t.Log(parse1(tt))
}

func parse1(token string) bool {
	now := int32(time.Now().Unix())
	t, err := jwt.Parse(token, nil)
	if t == nil {
		logx.Errorf("parseJwtToken, invalid token: %s, err: %+v", token, err)
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		logx.Infof("userId: %+v", claims["userId"])
		logx.Infof("expire_time: %+v", claims["expire_time"])
		if claims["expire_time"].(float64) < float64(now) {
			return false
		}
		// todo 续期
	}
	return true
}
