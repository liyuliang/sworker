package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/sworker/system"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/configmodel"
	"github.com/BurntSushi/toml"
	"os"
	"github.com/liyuliang/sworker/worker"
	"log"
)

func Start(port string) {

	r := gin.Default()
	r.GET("/profile", profile)
	r.GET("/spider", spider)
	r.GET("/pre", pre)

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

	u := "https://www.gufengmh8.com/manhua/bailianchengshen/"
	tpl := system.Config()["gufengmh8_list"]

	model := new(configmodel.Actions)
	_, err := toml.Decode(tpl, model)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	for i, a := range model.Action {
		if a.Target.Key == "ur" && a.Operation.Type == "download" {
			model.Action[i].Target.Value = u
		}

		a.Target.Value = u
		worker.Run(a)
	}

	data := worker.ReturnData()

	data.ToUrlVals()
	c.JSON(200, data)

}

func pre(c *gin.Context) {

	u := c.GetString("url")
	t := c.GetString("type")
	tpl := system.Config()[t]
	if tpl == "" {
		c.String(200, "invalid params")
		return
	}

	model := new(configmodel.Actions)
	_, err := toml.Decode(tpl, model)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	log.Print(len(model.Action))

	for i, a := range model.Action {
		if a.Target.Key == "ur" && a.Operation.Type == "download" {
			model.Action[i].Target.Value = u
		}

		a.Target.Value = u
		worker.Run(a)
	}

	c.JSON(200, worker.ReturnData())

}
