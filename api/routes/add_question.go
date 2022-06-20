package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/question/api/models"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func (store *Store) createQuestion(ctx *gin.Context) {
	var data models.Question
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = store.Querier.AddQuestion(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}
