package setups

import "github.com/gin-gonic/gin"

func LoadFiles(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("frontend/html/**/**/*")

	r.Static("/frontend/css/components", "./frontend/css/components")
	r.Static("/frontend/css/pages", "./frontend/css/pages")
	r.Static("/frontend/assets", "./frontend/assets")

	r.Static("/frontend/js/components", "./frontend/js/components")
	r.Static("/frontend/js/pages", "./frontend/js/pages")

	r.Static("/frontend/js/helpers", "./frontend/js/helpers")

	return r
}
