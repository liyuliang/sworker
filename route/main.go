package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/sworker/system"
	"github.com/liyuliang/utils/format"
)

func Start(port string) {

	r := gin.Default()
	r.GET("/profile", profile)
	r.GET("/spider", spider)

	r.NoRoute(method404)
	r.Run(":" + port)
}

func profile(c *gin.Context) {

	conf := system.Config()
	data := make(map[string]string)
	if len(conf) > 0 {

		data["system"] = conf["system"]
		data["core"] = conf["core"]
		data["load"] = conf["load"]
		data["memory"] = conf["memory"]
		data["disk"] = conf["disk"]

	} else {

		data["system"] = system.GetLinuxVersion()
		data["core"] = format.IntToStr(system.GetCoreNum())
		data["load"] = system.GetLoadAverage()
		data["memory"] = system.GetMemUsage()
		data["disk"] = system.GetDiskUsage()
	}
	c.JSON(200, data)
}

func spider(c *gin.Context) {

}
