package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/question/api/models"
	"github.com/gin-gonic/gin"
)

func (s *Store) signIn(c *gin.Context) {
	var data models.SignIn

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	login, err := s.Querier.SignIn(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	signInData, err := s.TokenMaker.CreateToken(login, s.Config.AccessDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	expTime := s.Config.AccessDuration.Seconds()
	c.SetCookie("token", signInData.Token, int(expTime), "/", s.Config.Host, false, true)
}

func (s *Store) welcome(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	claims, err := s.TokenMaker.VerifyToken(cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("Welcome %s!", claims.Login))
}
