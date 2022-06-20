package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func (store *Store) getUserInfo(ctx *gin.Context) {
	sort := ctx.Request.URL.Query().Get("sort")
	nickname := ctx.Param("user")

	userInfo, err := store.Querier.GetUserInfo(ctx, nickname, sort)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}
