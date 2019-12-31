package worker

import (
	"github.com/liyuliang/configmodel"
	"strings"
)

func init() {
	Register(func() worker {
		return new(trim)
	})
}

type trim struct {
}

func (d *trim) Name() string {
	return "trim"
}

func (d *trim) Do(a configmodel.Action) {

	t := getTempTarget(a)
	if t != nil {
		v, ok := t.(string)

		if ok {

			v = strings.Trim(v, a.Target.Value)
			if v != "" {
				replaceTempData(a.Target.Key, v)
			}
		}
	}
}
