package system

import (
	"github.com/liyuliang/utils/request"
	"encoding/json"
	"encoding/base64"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/utils/regex"
	"log"
)

var c format.MapData

func init() {
	c = format.Map()
}
func Config() format.MapData {
	return c
}

func Init(data format.MapData) {
	c = data

	initSystemUsed()
	initSpiderConfig(data)
}

func initSpiderConfig(data format.MapData) {
	gateway := data["gateway"]

	if gateway == "" {
		panic("gateway is required")
	}

	authApi := gateway + AuthApiPath

	resp, err := request.HttpPost(authApi, c.ToUrlVals())
	if err != nil {
		panic(err)
	}

	token := regex.Get(resp, `"uuid":"([^\"]+)"`)
	c["token"] = token

	log.Print(token)

	tplsApi := gateway + TplApiPath
	tpls, err := request.HttpPost(tplsApi, c.ToUrlVals())

	log.Print(tpls)

	model := make(map[string]string)
	json.Unmarshal([]byte(tpls), &model)
	for key, value := range model {
		c[key] = string(value)
		log.Print(key)
		continue
		v, err := base64.StdEncoding.DecodeString(value)
		if err == nil {
			c[key] = string(v)

			log.Print(key)
		} else {
			panic(err)
		}
	}
}

func initSystemUsed() {

	c["host"] = GetHostName()
	c["system"] = GetLinuxVersion()
	c["core"] = format.IntToStr(GetCoreNum())
	c["load"] = GetLoadAverage()
	c["memory"] = GetMemUsage()
	c["disk"] = GetDiskUsage()
}
