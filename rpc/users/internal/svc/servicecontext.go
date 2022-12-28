package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zhangphh_users/rpc/users/internal/config"
	"zhangphh_users/rpc/users/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserTabModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserTabModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
