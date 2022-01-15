package handler

import (
	"net/http"

	"go-zero-demo/greet/internal/logic"
	"go-zero-demo/greet/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func HomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHomeLogic(r.Context(), svcCtx)
		err := l.Home(w)
		if err != nil {
			httpx.Error(w, err)
		}
	}
}
