package post

import (
	"encoding/json"
	"serfwerk/server/misc"
	"serfwerk/server/misc/classes"
	"serfwerk/server/misc/cookie"
	"serfwerk/server/readers"

	"github.com/gin-gonic/gin"
)

func PostLogin(ctx *gin.Context) {
	var login_req classes.LoginReq
	ctx.BindJSON(&login_req)
	err := misc.CheckReqxpLogin(login_req.Username, login_req.Password)
	if err.Exit {
		classes.CreateResponse(ctx, err.Code, string(classes.Status), err.ErrString)
		return
	}
	cookie_val, err2 := readers.ReadLogin(login_req)
	if err2.Exit {
		classes.CreateResponse(ctx, err2.Code, string(classes.Status), err2.ErrString)
		return
	}
	cookie_head := cookie.CreateCookieLogin(cookie_val)
	ctx.Header("Set-Cookie", cookie_head)

	redirect_class := classes.RedirectHolder{Url: "/user", Msg: "Succes"}
	redirect_b, err3 := json.Marshal(redirect_class)
	if err3 != nil {
		classes.CreateResponse(ctx, 500, string(classes.Status), "Fejl under dekryption")
		return
	}
	str_redirect := string(redirect_b)
	classes.CreateResponse(ctx, 200, string(classes.Redirect), str_redirect)
}
