package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func (store *Store) getTodayAnswers(ctx *gin.Context) {
	page := ctx.Request.URL.Query().Get("page")

	todaysInfo, err := store.Querier.GetTodayAnswers(ctx, page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todaysInfo)
}
