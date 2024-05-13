package classes

import "github.com/gin-gonic/gin"

type UserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RespTypes string

const (
	Status   RespTypes = "status"
	Msg      RespTypes = "msg"
	Redirect RespTypes = "redirect"
)

type RedirectHolder struct {
	Url string `json:"url"`
	Msg string `json:"msg"`
}

type CustResponse struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func CreateResponse(ctx *gin.Context, code int, i_type string, msg string) {
	cust_rep := CustResponse{Code: code, Type: i_type, Msg: msg}
	ctx.JSON(code, cust_rep)
}
