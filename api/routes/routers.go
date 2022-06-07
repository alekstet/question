package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(s Store) *gin.Engine {
	s.Register()
	return s.Routes
}

/* func (rr *R) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rr.Router.ServeHTTP(w, r)
} */

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
