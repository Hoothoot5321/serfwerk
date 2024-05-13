package vm

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"serfwerk/server/cloudf"
	"serfwerk/server/misc"
	"time"
)

func NewAppVm(app_name string, username string, filename string, app_num int, port_num int, ram int, cpu int, cpu_kraft int, lager int) {
	ip := misc.IP
	misc.ExecPWSH("vboxm", "clonevm", "ubuntest8", "--name="+app_name, "--register")
	misc.ExecPWSH("vboxm", "modifyvm", app_name, "--natpf1", "delete", "base")
	misc.ExecPWSH("vboxm", "modifyvm", app_name, "--natpf1", fmt.Sprintf("\"base,tcp,%s,%d,,%d\"", ip, app_num, port_num))
	misc.ExecPWSH("vboxm", "modifyvm", app_name, "--natpf1", "delete", "helper")
	misc.ExecPWSH("vboxm", "modifyvm", app_name, "--natpf1", fmt.Sprintf("\"helper,tcp,%s,%d,,%d\"", ip, app_num+1, 9999))
	misc.ExecPWSH("vboxm", "modifyvm", app_name, "--natpf1", "delete", "shh")
	misc.ExecPWSH("vboxm", "modifyvm", app_name, "--natpf1", fmt.Sprintf("\"shh,tcp,%s,%d,,%d\"", ip, app_num+2, 22))
	misc.ExecPWSH("vboxm", "sharedfolder", "remove", app_name, "--name=shared")
	misc.ExecPWSH("mkdir", fmt.Sprintf("\"%s%s\\shared\"", misc.VBPATH, app_name))
	misc.ExecPWSH("vboxm", "sharedfolder", "add", app_name, "--name=shared", "--hostpath="+fmt.Sprintf("\"%s%s\\shared\"", misc.VBPATH, app_name), "--readonly", "--automount")
	misc.ExecPWSH("vboxm", "modifyvm", app_name, fmt.Sprintf("--memory=%d", ram), fmt.Sprintf("--cpus=%d", cpu), fmt.Sprintf("--cpuexecutioncap=%d", cpu_kraft))
	misc.ExecPWSH("vboxm", "modifymedium", fmt.Sprintf("\"%s%s\\%s.vdi\"", misc.VBPATH, app_name, app_name), fmt.Sprintf("--resize=%d", lager))
	misc.ExecPWSH("Copy-Item", "-Path", "lin_go", "-Destination", fmt.Sprintf("\"%s%s\\shared\"", misc.VBPATH, app_name))
	misc.ExecPWSH("Copy-Item", "-Path", "db/users/"+username+"/"+app_name+"/"+filename, "-Destination", fmt.Sprintf("\"%s%s\\shared\"", misc.VBPATH, app_name))

	os.WriteFile(fmt.Sprintf("%s%s\\shared\\gb.txt", misc.VBPATH, app_name), []byte(fmt.Sprintf("%dMB", lager)), 0644)

	misc.ExecPWSH("vboxm", "startvm", app_name, "--type=headless")
	time.Sleep(time.Millisecond * 500)
	go func() {
		cmd1 := exec.Command("pwsh", "-Command", "echo", "y")
		cmd2 := exec.Command("pwsh", "-Command", "plink", "-ssh", misc.USER+"@"+ip, "-P", fmt.Sprintf("%d", app_num+2), "-pw", misc.PSWD)
		r, w := io.Pipe()
		cmd1.Stdout = w
		cmd2.Stdin = r
		cmd2.Stdout = misc.CustWriter{}
		cmd1.Start()
		cmd2.Start()
		cmd1.Wait()
		w.Close()
		cmd2.Wait()
	}()
	go func() {
		time.Sleep(time.Second * 20)
		tunnel_id := cloudf.CreateTunnel(app_name)
		cloudf.CreateTunnelFile(app_name, username, tunnel_id, app_num)
		cloudf.RouteTunnel(tunnel_id, app_name)
		cloudf.RunTunnel(username, app_name)
	}()

}
