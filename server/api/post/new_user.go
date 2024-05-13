package post

import (
	"encoding/json"
	"serfwerk/server/misc"
	"serfwerk/server/misc/classes"
	"serfwerk/server/misc/cookie"
	"serfwerk/server/writers"

	"github.com/gin-gonic/gin"
)

func PostNewUser(ctx *gin.Context) {
	var user_req classes.UserReq
	ctx.BindJSON(&user_req)
	cust_resp := misc.CheckReqxp(user_req.Username, user_req.Email, user_req.Password)
	if cust_resp.Exit {
		classes.CreateResponse(ctx, cust_resp.Code, string(classes.Status), cust_resp.ErrString)
		return
	}
	redirect_holder := classes.RedirectHolder{Url: "/user", Msg: "Succes"}
	redirect_msg, err := json.Marshal(redirect_holder)
	if err != nil {
		classes.CreateResponse(ctx, 500, string(classes.Status), "Fejl under dekryption")
		return
	}
	str_redirect_msg := string(redirect_msg)
	cookie_val, err2 := writers.WriteNewUser(user_req)
	if err2.Exit {
		classes.CreateResponse(ctx, err2.Code, string(classes.Status), err2.ErrString)
		return
	}
	ctx.Header("Set-Cookie", cookie.CreateCookieLogin(cookie_val))
	classes.CreateResponse(ctx, 200, string(classes.Redirect), str_redirect_msg)
	return
}
