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
	// todo: add your logic here and delete this line

	return
}
