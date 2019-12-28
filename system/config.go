package system

import (
	"github.com/liyuliang/utils/request"
	"encoding/json"
	"encoding/base64"
	"github.com/liyuliang/utils/format"
	"log"
	"os"
)

type appConfig map[string]string

var _config appConfig

func init() {
	_config = make(map[string]string)
}
func Config() appConfig {
	return _config
}

func Init(gateway, auth string) {

	initSystemUsed()
	initSpiderConfig(gateway, auth)
}

func initSpiderConfig(gateway string, auth string) {
	//gateway
	resp := request.HttpGet(gateway)
	if resp.Err != nil {
		panic(resp.Err)
	}
	//if !strings.Contains(resp.Data, "success") {
	//fmt.Fprintf(os.Stderr, "Auth failed \n")
	//os.Exit(2)
	//}

	model := make(map[string]string)
	json.Unmarshal([]byte(resp.Data), model)
	for key, value := range model {
		v, err := base64.StdEncoding.DecodeString(value)
		if err == nil {
			_config[key] = string(v)
		} else {
			panic(err)
		}
	}
}

func initSystemUsed() {

	_config["system"] = GetLinuxVersion()
	_config["core"] = format.IntToStr(GetCoreNum())
	_config["load"] = GetLoadAverage()
	_config["memory"] = GetMemUsage()
	_config["disk"] = GetDiskUsage()
}
