package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/question/api/models"
	"github.com/gin-gonic/gin"
)

func (store *Store) signIn(ctx *gin.Context) {
	var data models.SignIn

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	login, err := store.Querier.SignIn(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	signInData, err := store.TokenMaker.CreateToken(login, store.Config.AccessDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	expTime := store.Config.AccessDuration.Seconds()
	ctx.SetCookie("token", signInData.Token, int(expTime), "/", store.Config.Host, false, true)
}

func (store *Store) welcome(ctx *gin.Context) {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	claims, err := store.TokenMaker.VerifyToken(cookie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("Welcome %s!", claims.Login))
}
