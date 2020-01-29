package worker

import (
	"github.com/liyuliang/configmodel"
	"github.com/liyuliang/sworker/system"
	"strings"
)

func init() {
	Register(func() worker {
		return new(array_prefix)
	})
}

type array_prefix struct {
}

func (d *array_prefix) Name() string {
	return "array_prefix"
}

func (d *array_prefix) Do(a configmodel.Action) {

	var data interface{}

	if a.Operation.Option.Type == system.ActionTempPoolName {
		data = getTempData(a.Operation.Option.Key)
	}

	if data == nil {
		return
	}

	prefix := data.(string)

	t := getTempTarget(a)

	var newArr []string

	if t != nil {
		v, ok := t.([]string)
		if ok {

			for _, value := range v {
				if !strings.Contains(value, prefix) {
					value = prefix + value
				}
				newArr = append(newArr, value)
			}
			replaceTempData(a.Target.Key, newArr)
		}
	}
	return
}
