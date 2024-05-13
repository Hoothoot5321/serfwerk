package setups

import (
	"fmt"
	"os"
	"serfwerk/server/misc/classes"
	"serfwerk/server/readers"
	"strings"

	"github.com/gin-gonic/gin"
)

func getLang(ctx *gin.Context) string {
	lang_header := ctx.Request.Header.Get("Accept-Language")
	lang := strings.Split(lang_header, ",")[0]
	return lang
}

func getBaseName(page_name string) string {
	base_name := strings.Split(page_name, ".")[0]
	base_name = strings.ReplaceAll(base_name, "-", "/")
	return base_name
}

func showPage(ctx *gin.Context, lang string, page_name string) {
	if lang != "å" {
		fmt.Println("Danish")
		fmt.Println(page_name)
		ctx.HTML(200, page_name, gin.H{})
	} else {
		fmt.Println("English")
		ctx.HTML(200, "en_"+page_name, gin.H{})
	}
}

func showUserPage(ctx *gin.Context, lang string, page_name string, user classes.User) {
	if lang != "å" {
		fmt.Println("Danish")
		fmt.Println(page_name)
		ctx.HTML(200, page_name, gin.H{
			"User": user,
			"Apps": user.Apps,
		})
	} else {
		fmt.Println("English")
		ctx.HTML(200, "en_"+page_name, gin.H{})
	}
}

func SetupHTML(r *gin.Engine) *gin.Engine {
	base_html_path := "frontend/html/da/"

	html_pages, _ := os.ReadDir(base_html_path + "pages")

	login_pages, _ := os.ReadDir(base_html_path + "login_pages")

	user_pages, _ := os.ReadDir(base_html_path + "user_pages")

	for _, page := range html_pages {
		base_name := getBaseName(page.Name())
		if base_name == "index" {
			base_name = "/"
		} else {
			base_name = "/" + base_name
		}
		func(base_name2 string, page_name string) {
			r.GET(base_name2, func(ctx *gin.Context) {
				lang := getLang(ctx)
				showPage(ctx, lang, page_name)
			})
		}(base_name, page.Name())
	}
	for _, page := range login_pages {
		base_name := getBaseName(page.Name())
		base_name = "/" + base_name
		func(base_name2 string, page_name string) {
			r.GET(base_name2, func(ctx *gin.Context) {
				lang := getLang(ctx)
				cookie_val, err := ctx.Cookie("auth_cookie")
				if err == nil {
					_, cust_err := readers.ReadCookie(cookie_val)
					if cust_err.Exit {
						showPage(ctx, lang, page_name)
					} else {
						ctx.Redirect(303, "/user")
					}
				} else {
					showPage(ctx, lang, page_name)
				}
			})
		}(base_name, page.Name())
	}
	for _, page := range user_pages {
		base_name := getBaseName(page.Name())
		if base_name == "appboard" {
			base_name = "/"
		} else {
			base_name = "/" + base_name
		}
		func(base_name2 string, page_name string) {
			r.GET("/user"+base_name2, func(ctx *gin.Context) {
				lang := getLang(ctx)
				cookie_val, err := ctx.Cookie("auth_cookie")
				if err == nil {
					user, cust_err := readers.ReadCookie(cookie_val)
					if cust_err.Exit {
						ctx.Redirect(303, "/login")
					} else {
						showUserPage(ctx, lang, page_name, user)
					}
				} else {

					ctx.Redirect(303, "/login")
				}
			})
		}(base_name, page.Name())
	}
	return r
}
