package worker

import (
	"github.com/liyuliang/configmodel"
	"strings"
)

func init() {
	Register(func() worker {
		return new(strings_replace)
	})
}

type strings_replace struct {
}

func (d *strings_replace) Name() string {
	return "strings_replace"
}

func (d *strings_replace) Do(a configmodel.Action) {

	t := getTempTarget(a)
	if t != nil {
		v, ok := t.(string)

		if ok {

			v = strings.Replace(v, a.Target.Value, a.Operation.Value, -1)
			if v != "" {
				replaceTempData(a.Target.Key, v)
			}
		}
	}
}
