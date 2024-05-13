package cloudf

import (
	"fmt"
	"os"
	"os/exec"
	"serfwerk/server/misc"
	"strings"
)

func CreateTunnel(tunnel_name string) string {
	asdf, _ := exec.Command("pwsh", "-Command", "C:\\Users\\MartinNammat\\Documents\\Programming-2\\Projects\\Serfwerk2\\cloudflared.exe", "tunnel", "create", tunnel_name).Output()
	s_asd := string(asdf)
	fmt.Println(string(s_asd))
	ind := strings.Index(s_asd, "id")
	id_test := s_asd[ind+3:]
	id_test = strings.TrimSpace(id_test)
	return id_test
}

func CreateTunnelFile(appname string, username string, tunnel_id string, port_num int) {
	db_path := misc.DBPath
	in_str := fmt.Sprintf("tunnel: %s\ncredentials-file: C:\\Users\\MartinNammat\\.cloudflared\\%s.json\ningress:\n- hostname: %s.serfwerk.online\n  service: http://192.168.1.12:%d\n- service: http_status:404", tunnel_id, tunnel_id, appname, port_num)
	os.WriteFile(db_path+"users/"+username+"/"+appname+"/config.yml", []byte(in_str), 0644)
}
func RouteTunnel(tunnel_id string, appname string) {
	ou, _ := exec.Command("pwsh", "-Command", "C:\\Users\\MartinNammat\\Documents\\Programming-2\\Projects\\Serfwerk2\\cloudflared.exe", "tunnel", "route", "dns", tunnel_id, appname).Output()
	fmt.Println(string(ou))
}
func RunTunnel(username string, appname string) {
	ou, _ := exec.Command("pwsh", "-Command", "C:\\Users\\MartinNammat\\Documents\\Programming-2\\Projects\\Serfwerk2\\cloudflared.exe", "tunnel", "--config", fmt.Sprintf("C:\\Users\\MartinNammat\\Documents\\Programming-2\\Projects\\Serfwerk2\\db\\users\\%s\\%s\\config.yml", username, appname), "run", appname).Output()
	fmt.Println(string(ou))
	fmt.Println("FUUUUUUUUUC")
}
