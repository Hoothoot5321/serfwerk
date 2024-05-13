package main

import (
	"serfwerk/server/misc"
	"serfwerk/server/setups"

	"github.com/gin-gonic/gin"
)

func main() {
	setups.SetupDB()
	r := gin.Default()
	r.MaxMultipartMemory = 100 << 20

	r = setups.LoadFiles(r)
	r = setups.SetupHTML(r)
	r = setups.SetupPOST(r)

	go func() {
		misc.ExecPWSH("cloudflared", "tunnel", "--config", "serfwerk_config.yml", "run", "serfwerk")
	}()
	r.Run()
}
