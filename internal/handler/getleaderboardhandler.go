package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hackathon_be/internal/logic"
	"hackathon_be/internal/svc"
)

func getleaderboardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetleaderboardLogic(r.Context(), svcCtx)
		resp, err := l.Getleaderboard()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
