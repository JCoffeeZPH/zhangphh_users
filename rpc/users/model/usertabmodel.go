package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserTabModel = (*customUserTabModel)(nil)

type (
	// UserTabModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTabModel.
	UserTabModel interface {
		userTabModel
	}

	customUserTabModel struct {
		*defaultUserTabModel
	}
)

// NewUserTabModel returns a model for the database table.
func NewUserTabModel(conn sqlx.SqlConn, c cache.CacheConf) UserTabModel {
	return &customUserTabModel{
		defaultUserTabModel: newUserTabModel(conn, c),
	}
}
