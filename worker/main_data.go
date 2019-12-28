package worker


import (
	"github.com/liyuliang/configmodel"
)


var tempData map[string]interface{}
var returnData map[string]interface{}

func init() {
	tempData = make(map[string]interface{})
	returnData = make(map[string]interface{})
}

func TempData() map[string]interface{} {
	return tempData
}

func ReturnData() map[string]interface{} {
	return returnData
}

func setTempData(k string, v interface{}) {
	_, exist := tempData[k]
	if exist {
		return
	} else {
		tempData[k] = v
	}
}
func getTempData(k string) interface{} {

	if k == "" {
		return nil
	}
	_, exist := tempData[k]
	if exist {
		return tempData[k]
	}
	return nil
}

func getTempTarget(a configmodel.Action) interface{} {
	if a.Target.Type == "temp" {
		return tempData[a.Target.Key]
	}
	return nil
}

func replaceTempData(k string, v interface{}) {
	tempData[k] = v
}

func setReturnData(k string, v interface{}) {
	_, exist := returnData[k]
	if exist {
		return
	} else {
		returnData[k] = v
	}
}
func replaceReturnData(k string, v interface{}) {
	returnData[k] = v
}
