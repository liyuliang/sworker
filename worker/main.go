package worker

import (
	"github.com/liyuliang/configmodel"
	"sync"
	"strings"
	"log"
)

var tempData map[string]interface{}
var returnData map[string]interface{}

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

func init() {
	tempData = make(map[string]interface{})
	returnData = make(map[string]interface{})
}

func Run(a configmodel.Action) {

	if a.Target.Value != "" {
		setTempData(a.Target.Key, a.Target.Value)
	}

	w := Get(a.Operation.Type)
	if w == nil {
		log.Printf("worker %s is missing from operation type", a.Operation.Type)
		return
	}
	//log.Printf("worker %s prepare run", w().Name())
	data := w().Do(a)

	if data != "" {
		setReturnData(a.Return, data)
	}

}

type worker interface {
	Name() string
	Do(a configmodel.Action) string
}

type Creator func() worker

var _list []Creator

func Register(method Creator) {
	_list = append(_list, method)
}

type parserList struct {
	sync.RWMutex
	creators map[string]Creator
}

var list parserList

func List() map[string]Creator {

	if len(list.creators) != len(_list) {

		list = parserList{}
		list.creators = make(map[string]Creator)
		list.Lock()

		for _, agent := range _list {
			list.creators[agent().Name()] = agent
		}
		list.Unlock()
	}

	return list.creators
}

func Get(name string) (creator Creator) {
	for _, agent := range List() {
		if strings.ToLower(agent().Name()) == strings.ToLower(name) {
			creator = agent
			break
		}
	}
	return creator
}
