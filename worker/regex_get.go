package worker

import (
	"github.com/liyuliang/configmodel"
	"github.com/liyuliang/utils/regex"
)

func init() {
	Register(func() worker {
		return new(regex_get)
	})
}

type regex_get struct {
}

func (d *regex_get) Name() string {
	return "regex_get"
}

func (d *regex_get) Do(a configmodel.Action) string {
	t := getTempTarget(a)
	if t != nil {
		v, ok := t.(string)

		if ok {
			newV := regex.Get(v, a.Operation.Value)
			if newV != "" {
				setTempData(a.Target.Key, newV)
				setTempData(a.Operation.Key, newV)
			}
		}
	}
	return ""
}
