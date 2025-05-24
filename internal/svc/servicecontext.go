package svc

import (
	"hackathon_be/internal/config"
	"hackathon_be/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	ScoreModel model.ScoreModel
}

func NewServiceContext(c config.Config, conn sqlx.SqlConn) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ScoreModel: model.NewScoreModel(conn),
	}
}

