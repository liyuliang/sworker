package system

import (
	"github.com/liyuliang/utils/request"
	"encoding/json"
	"encoding/base64"
	"github.com/liyuliang/utils/format"
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

	//gateway
	token, err := request.HttpPost(gateway, c.ToUrlVals())
	if err != nil {
		panic(err)
	}

	model := make(map[string]string)
	json.Unmarshal([]byte(resp.Data), model)
	for key, value := range model {
		v, err := base64.StdEncoding.DecodeString(value)
		if err == nil {
			c[key] = string(v)
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
