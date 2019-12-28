package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/sworker/system"
)

func Start(port string) {

	r := gin.Default()
	r.GET("/profile", profile)

	r.NoRoute(method404)
	r.Run(":8888")
}

func profile(c *gin.Context) {

	data := make(map[string]string)
	data["system"] = system.GetLinuxVersion()
	data["core"] = format.IntToStr(system.GetCoreNum())
	data["load"] = system.GetLoadAverage()
	data["memory"] = system.GetMemUsage()
	data["disk"] = system.GetDiskUsage()
	c.JSON(200, data)
}
