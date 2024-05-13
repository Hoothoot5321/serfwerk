package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var usernames [3]string
	usernames[0] = "Martin"
	usernames[1] = "Jhon"
	usernames[2] = "Jens"
	r.Static("/assets/js", "./js")
	r.LoadHTMLGlob("./html/**/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "pages/loading.html", gin.H{})
	})
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "pages/loading.html", gin.H{})
	})

	r.GET("/pre/", func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		ctx.HTML(http.StatusOK, "pages/index.html", gin.H{
			"username": token,
			"temp":     usernames,
			"fish":     "Hello Jhonny!!!",
		})
	})
	r.GET("/test/", func(ctx *gin.Context) {
		fmt.Println("Received request")
		ctx.String(200, "Hello boy!!!")
	})
	r.POST("/api/bang", func(ctx *gin.Context) {
		post_num := ctx.DefaultPostForm("portnum", "none")
		fmt.Println(post_num)
	})

	r.Run(":7000")
}
