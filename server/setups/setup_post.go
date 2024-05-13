package setups

import (
	"serfwerk/server/api/post"

	"github.com/gin-gonic/gin"
)

func SetupPOST(r *gin.Engine) *gin.Engine {
	r.POST("/api/new_user", post.PostNewUser)

	r.POST("/api/login", post.PostLogin)
	r.POST("/api/new_app", post.PostNewApp)

	return r
}
