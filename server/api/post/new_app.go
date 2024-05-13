package post

import (
	"encoding/json"
	"fmt"
	"os"
	"serfwerk/server/misc"
	"serfwerk/server/misc/classes"
	"serfwerk/server/vm"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
)

func PostNewApp(ctx *gin.Context) {
	fmt.Println("Got msg")

	db_path := misc.DBPath
	file, _ := ctx.FormFile("file")
	app_name := ctx.DefaultPostForm("appName", "")
	run_cmd := ctx.DefaultPostForm("runCmd", "")
	build_cmd := ctx.DefaultPostForm("buildCmd", "")
	port_num, _ := strconv.Atoi(ctx.DefaultPostForm("portNum", ""))
	cookie_val, _ := ctx.Cookie("auth_cookie")
	ram, _ := strconv.Atoi(ctx.DefaultPostForm("ram", ""))
	pris, _ := strconv.Atoi(ctx.DefaultPostForm("pris", ""))
	fmt.Println(pris)
	lager, _ := strconv.Atoi(ctx.DefaultPostForm("lager", ""))
	cpu, _ := strconv.Atoi(ctx.DefaultPostForm("cpu", ""))
	cpu_kraft, _ := strconv.Atoi(ctx.DefaultPostForm("cpu_kraft", ""))

	b_content, _ := os.ReadFile(db_path + "users.json")
	var user_list []classes.User
	json.Unmarshal(b_content, &user_list)
	var user classes.User
	found_match := false
	user_ind := 0
	app_num_b, _ := os.ReadFile(db_path + "app_num.txt")
	fmt.Println("App nume")
	app_num_b_s := strings.Trim(string(app_num_b), "\n\r ")

	fmt.Println(app_num_b_s)
	app_num, err := strconv.Atoi(app_num_b_s)
	if err != nil {
		fmt.Println("Shit")
	}

	fmt.Println(app_num)
	for i, n_user := range user_list {
		if n_user.UserCookie == cookie_val {
			user = n_user
			found_match = true
			user_ind = i
		}
		for _, app := range n_user.Apps {
			app_num += 3
			if app.AppName == app_name {
				classes.CreateResponse(ctx, 400, string(classes.Status), "App navn eksisterer allerede")
				return
			}
		}
	}
	app_num += 3
	s_app_num := strconv.Itoa(app_num)
	fmt.Println(s_app_num)
	os.Remove(db_path + "app_num.txt")
	os.WriteFile(db_path+"app_num.txt", []byte(s_app_num), 0644)

	if !found_match {
		classes.CreateResponse(ctx, 400, string(classes.Status), "Invalid cookie")
		return
	}
	fmt.Printf("AppName: %s\nRunCmd: %s\nBuildCmd: %s\nCookieVal: %s\nPortNum: %d\n", app_name, run_cmd, build_cmd, cookie_val, port_num)
	app_path := db_path + "users/" + user.Username + "/" + app_name
	file_path := app_path + "/" + file.Filename

	os.Mkdir(app_path, 0744)
	ctx.SaveUploadedFile(file, file_path)
	buf, err2 := os.ReadFile(file_path)
	if err2 != nil {
		classes.CreateResponse(ctx, 500, string(classes.Status), "Fejl i at tjekke fil type")
		return
	}
	f_type, err2 := filetype.Match(buf)
	if err2 != nil {
		classes.CreateResponse(ctx, 500, string(classes.Status), "Fejl i at tjekke fil type")
		os.RemoveAll(app_path)
		return
	}

	if f_type.MIME.Subtype != "zip" {
		classes.CreateResponse(ctx, 400, string(classes.Status), "Uploadet fil er ikke zip fil")
		os.RemoveAll(app_path)
		return
	}
	new_app := classes.App{AppName: app_name, OwnerCookie: user.UserCookie, RunCmd: run_cmd, BuilCmd: build_cmd, PortNum: port_num, Ram: ram, CPU: cpu, CPUPower: cpu_kraft, Lager: lager, AppNum: app_num}
	user.Apps = append(user.Apps, new_app)

	user_list[user_ind] = user

	user_byte, _ := json.Marshal(user_list)
	os.WriteFile(db_path+"users.json", user_byte, 0644)
	vm.NewAppVm(app_name, user.Username, file.Filename, app_num, port_num, ram, cpu, cpu_kraft, lager+8192)
	classes.CreateResponse(ctx, 200, string(classes.Status), "Nice")

	return
}
