package worker

import "github.com/liyuliang/configmodel"

type array_prefix struct {
}

func (d *array_prefix) Name() string {
	return "array_prefix"
}

func (d *array_prefix) Do(a configmodel.Action) string{

	t := getTempTarget(a)
	if t != nil {
		v, ok := t.([]string)
		if ok {
			for _, value := range v {

			}
		}
	}
}
