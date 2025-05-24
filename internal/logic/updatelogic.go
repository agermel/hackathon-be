package logic

import (
	"context"

	"hackathon_be/internal/model"
	"hackathon_be/internal/svc"
	"hackathon_be/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.ScoreReq) (resp *types.GeneralRes, err error) {
	var score model.Score
	score.Score = int64(req.Score)
	score.Username = req.Username

	err = l.svcCtx.ScoreModel.Update(l.ctx, &score)
	if err != nil {
		return nil, err
	}

	return &types.GeneralRes{
		Base: types.Base{
			Code: 200,
			Msg:  "success",
		},
		Data: types.EmptyData{},
	}, err
}
