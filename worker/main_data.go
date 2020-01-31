package worker

import (
	"github.com/liyuliang/configmodel"
	"github.com/liyuliang/sworker/system"
	"github.com/liyuliang/utils/format"
	"reflect"
	"net/url"
)

type Data map[string]interface{}

func (d Data) ToUrlVals() url.Values {

	vals := url.Values{}
	for k, v := range d {

		to := reflect.TypeOf(v)
		if to == nil {
			continue
		}

		t := to.String()
		println(k, t)

		switch t {
		case "string":
			vals.Add(k, v.(string))
		case "int":
			vals.Add(k, format.IntToStr(v.(int)))
		case "int64":
			vals.Add(k, format.Int64ToStr(v.(int64)))
		case "[]string":

			for _, value := range v.([]string) {
				vals.Add(k, value)
			}
		}
	}
	return vals
}

var tempData Data
var returnData Data
var statusCode int

func init() {
	Clean()
}

func Clean() {

	tempData = make(Data)
	returnData = make(Data)
	statusCode = 0
}

func TempData() Data {
	return tempData
}

func ReturnData() Data {
	return returnData
}
func StatusCode() int {
	return statusCode
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
	if a.Target.Type == system.ActionTempPoolName {
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
