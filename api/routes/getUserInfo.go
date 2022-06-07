package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) getUserInfo(c *gin.Context) {
	sort := c.Request.URL.Query().Get("sort")
	nickname := c.Param("user")

	userInfo, err := s.Querier.GetUserInfo(nickname, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
