package handler

import (
	"net/http"

	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/logic"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ListUserAddressByUserIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetByUserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListUserAddressByUserIdLogic(r.Context(), ctx)
		resp, err := l.ListUserAddressByUserId(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
