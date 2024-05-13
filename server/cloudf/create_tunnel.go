package cloudf

import (
	"fmt"
	"os"
	"os/exec"
	"serfwerk/server/misc"
	"strings"
)

func CreateTunnel(tunnel_name string) string {
	asdf, _ := exec.Command("pwsh", "-Command", "cloudflared", "tunnel", "create", tunnel_name).Output()
	s_asd := string(asdf)
	fmt.Println(string(s_asd))
	ind := strings.Index(s_asd, "id")
	id_test := s_asd[ind+3:]
	id_test = strings.TrimSpace(id_test)
	return id_test
}

func CreateTunnelFile(appname string, username string, tunnel_id string, port_num int) {
	db_path := misc.DBPath
	in_str := fmt.Sprintf("tunnel: %s\ncredentials-file: %s%s.json\ningress:\n- hostname: %s.serfwerk.online\n  service: http://%s:%d\n- service: http_status:404", tunnel_id, misc.CLPATH, tunnel_id, appname, misc.IP, port_num)
	os.WriteFile(db_path+"users/"+username+"/"+appname+"/config.yml", []byte(in_str), 0644)
}
func RouteTunnel(tunnel_id string, appname string) {
	misc.ExecPWSH("cloudflared", "tunnel", "route", "dns", tunnel_id, appname)
}
func RunTunnel(username string, appname string) {
	misc.ExecPWSH("cloudflared", "tunnel", "--config", fmt.Sprintf("db\\users\\%s\\%s\\config.yml", username, appname), "run", appname)
}
