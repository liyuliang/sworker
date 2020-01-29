package system

import (
	"github.com/liyuliang/utils/request"
	"encoding/json"
	"encoding/base64"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/utils/regex"
	"time"
)

var c format.MapData

func init() {
	c = format.Map()
}
func Config() format.MapData {
	return c
}

func SetConfig(key, val string) format.MapData {
	v := Config()
	v[key] = val
	return v
}

func Init(data format.MapData) {
	c = data

	initSystemUsed()
	initSpiderConfig(data)
}

func initSpiderConfig(data format.MapData) {
	gateway := data[SystemGateway]

	if gateway == "" {
		panic("gateway is required")
	}

	authApi := gateway + AuthApiPath

	resp, err := request.HttpPost(authApi, data.ToUrlVals())
	if err != nil {
		panic(err)
	}

	token := regex.Get(resp, `"uuid":"([^\"]+)"`)
	c[SystemToken] = token

	tplsApi := gateway + TplApiPath
	tpls, err := request.HttpPost(tplsApi, c.ToUrlVals())

	model := make(map[string]string)
	json.Unmarshal([]byte(tpls), &model)

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

	c["start"] = time.Now().String()
	c["host"] = GetHostName()
	c["system"] = GetLinuxVersion()
	c["core"] = format.IntToStr(GetCoreNum())
	c["load"] = GetLoadAverage()
	c["memory"] = GetMemUsage()
	c["disk"] = GetDiskUsage()
}
