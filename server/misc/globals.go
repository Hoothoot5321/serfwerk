package misc

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

var DBPath = "db/"
var IP = ""
var USER = ""
var PSWD = ""
var VBPATH = ""
var CLPATH = ""

func GetCurPath() string {
	cur_path, _ := os.Getwd()
	return cur_path
}

func GetALL() {
	b_c, _ := os.ReadFile("setting.json")
	var dict map[string]string
	json.Unmarshal(b_c, &dict)
	IP = dict["ip"]
	USER = dict["username"]
	PSWD = dict["password"]
	VBPATH = dict["vbox_path"]
	CLPATH = dict["cloud_cred"]
}

func ExecPWSH(args ...string) {
	cur_dir := GetCurPath()
	o := exec.Command("pwsh", "-command", "set-location", cur_dir, ";")
	o.Args = append(o.Args, args...)
	o.Stdout = CustWriter{}
	o.Run()

}

type CustWriter struct{}

func (e CustWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
