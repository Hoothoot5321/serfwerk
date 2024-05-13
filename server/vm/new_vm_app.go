package vm

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"serfwerk/server/cloudf"
	"time"
)

type custWriter struct{}

func (e custWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
func NewAppVm(app_name string, username string, filename string, app_num int, port_num int, ram int, cpu int, cpu_kraft int, lager int) {
	out1, _ := exec.Command("pwsh", "-Command", "vboxm", "clonevm", "ubuntest8", "--name="+app_name, "--register").Output()
	fmt.Println(string(out1))
	out2, _ := exec.Command("pwsh", "-Command", "vboxm", "modifyvm", app_name, "--natpf1", "delete", "base").Output()
	fmt.Println(string(out2))
	out3, _ := exec.Command("pwsh", "-Command", "vboxm", "modifyvm", app_name, "--natpf1", fmt.Sprintf("\"base,tcp,192.168.1.12,%d,,%d\"", app_num, port_num)).Output()

	fmt.Println(string(out3))
	out4, _ := exec.Command("pwsh", "-Command", "vboxm", "modifyvm", app_name, "--natpf1", "delete", "helper").Output()

	fmt.Println(string(out4))
	out5, _ := exec.Command("pwsh", "-Command", "vboxm", "modifyvm", app_name, "--natpf1", fmt.Sprintf("\"helper,tcp,192.168.1.12,%d,,%d\"", app_num+1, 9999)).Output()

	fmt.Println(string(out5))

	out11, _ := exec.Command("pwsh", "-Command", "vboxm", "modifyvm", app_name, "--natpf1", "delete", "shh").Output()
	fmt.Println(string(out11))

	out12, _ := exec.Command("pwsh", "-Command", "vboxm", "modifyvm", app_name, "--natpf1", fmt.Sprintf("\"shh,tcp,192.168.1.12,%d,,%d\"", app_num+2, 22)).Output()
	fmt.Println(string(out12))

	out6, _ := exec.Command("pwsh", "-Command", "vboxm", "sharedfolder", "remove", app_name, "--name=shared").Output()

	fmt.Println(string(out6))
	out7, _ := exec.Command("pwsh", "-Command", "mkdir", fmt.Sprintf("\"C:\\\\Users\\MartinNammat\\VirtualBox VMs\\%s\\shared\"", app_name)).Output()

	fmt.Println(string(out7))
	out8, _ := exec.Command("pwsh", "-Command", "vboxm", "sharedfolder", "add", app_name, "--name=shared", "--hostpath="+fmt.Sprintf("\"C:\\\\Users\\MartinNammat\\VirtualBox VMs\\%s\\shared\"", app_name), "--readonly", "--automount").Output()

	fmt.Println(string(out8))
	out9, _ := exec.Command("pwsh", "-Command", "vboxm", "modifyvm", app_name, fmt.Sprintf("--memory=%d", ram), fmt.Sprintf("--cpus=%d", cpu), fmt.Sprintf("--cpuexecutioncap=%d", cpu_kraft)).Output()

	fmt.Println(string(out9))
	out10, _ := exec.Command("pwsh", "-Command", "vboxm", "modifymedium", fmt.Sprintf("\"C:\\\\Users\\MartinNammat\\VirtualBox VMs\\%s\\%s.vdi\"", app_name, app_name), fmt.Sprintf("--resize=%d", lager)).Output()
	fmt.Println(string(out10))

	out13, _ := exec.Command("pwsh", "-Command", "Copy-Item", "-Path", "C:\\Users\\MartinNammat\\Documents\\Programming-2\\Projects\\Serfwerk2\\lin_go", "-Destination", fmt.Sprintf("\"C:\\\\Users\\MartinNammat\\VirtualBox VMs\\%s\\shared\"", app_name)).Output()
	fmt.Println(string(out13))

	out14, _ := exec.Command("pwsh", "-Command", "Copy-Item", "-Path", fmt.Sprintf("C:\\Users\\MartinNammat\\Documents\\Programming-2\\Projects\\Serfwerk2\\db\\users\\%s\\%s\\%s", username, app_name, filename), "-Destination", fmt.Sprintf("\"C:\\\\Users\\MartinNammat\\VirtualBox VMs\\%s\\shared\"", app_name)).Output()
	fmt.Println(string(out14))

	os.WriteFile(fmt.Sprintf("C:\\\\Users\\MartinNammat\\VirtualBox VMs\\%s\\shared\\gb.txt", app_name), []byte(fmt.Sprintf("%dMB", lager)), 0644)

	//putty -ssh hoothoot@192.168.1.12 8000 -pw hoothoot//

	out16, _ := exec.Command("pwsh", "-Command", "vboxm", "startvm", app_name, "--type=headless").Output()
	fmt.Println(string(out16))
	time.Sleep(time.Millisecond * 500)
	go func() {
		cmd1 := exec.Command("pwsh", "-Command", "echo", "y")
		cmd2 := exec.Command("pwsh", "-Command", "C:\\Users\\MartinNammat\\Documents\\Programming-2\\Projects\\Serfwerk2\\plink.exe", "-ssh", "hoothoot@192.168.1.12", "-P", fmt.Sprintf("%d", app_num+2), "-pw", "hoothoot")
		r, w := io.Pipe()
		cmd1.Stdout = w
		cmd2.Stdin = r
		cmd2.Stdout = custWriter{}
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
