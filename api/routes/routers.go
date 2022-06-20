package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(s Store) *gin.Engine {
	s.Register()
	return s.Routes
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
