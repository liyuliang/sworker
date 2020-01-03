package worker

import (
	"github.com/liyuliang/configmodel"
	"strings"
)

func init() {
	Register(func() worker {
		return new(strings_Split)
	})
}

type strings_Split struct {
}

func (d *strings_Split) Name() string {
	return "strings_split"
}

func (d *strings_Split) Do(a configmodel.Action)  {

	t := getTempTarget(a)
	if t != nil {

		v, ok := t.(string)
		if ok {
			vals := strings.Split(v, a.Operation.Value)
			setTempData(a.Operation.Key, vals)
		}
	}
	return
}
