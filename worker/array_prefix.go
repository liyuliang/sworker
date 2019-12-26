package worker

import (
	"github.com/liyuliang/configmodel"
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

	if a.Operation.Option.Type == "temp" {
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
				newV := prefix + value
				newArr = append(newArr, newV)
			}
			replaceTempData(a.Target.Key, newArr)
		}
	}
	return
}
