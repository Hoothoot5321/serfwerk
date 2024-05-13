package main

import (
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

	r.Run()
}
