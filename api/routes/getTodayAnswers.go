package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) getTodayAnswers(c *gin.Context) {
	page := c.Request.URL.Query().Get("page")

	todaysInfo, err := s.Querier.GetTodayAnswers(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, todaysInfo)
}
