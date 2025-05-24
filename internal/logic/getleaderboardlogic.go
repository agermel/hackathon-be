package logic

import (
	"context"

	"hackathon_be/internal/svc"
	"hackathon_be/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetleaderboardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetleaderboardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetleaderboardLogic {
	return &GetleaderboardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetleaderboardLogic) Getleaderboard() (resp *types.LeaderboardRes, err error) {
	scores, err := l.svcCtx.ScoreModel.FindTop10(l.ctx)
	if err != nil {
		return nil, err
	}

	var data []types.Score

	for _, score := range scores {
		data = append(data, types.Score{
			Username: score.Username,
			Score:    int(score.Score),
		})
	}

	return &types.LeaderboardRes{
		Base: types.Base{
			Code: 200,
			Msg:  "success",
		},
		Data: data,
	}, nil
}
