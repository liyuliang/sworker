package worker

import (
	"github.com/liyuliang/configmodel"
	"github.com/liyuliang/dom-parser"
)

func init() {

	Register(func() worker {
		return new(dom_find)
	})
}

type dom_find struct {
}

func (d *dom_find) Name() string {
	return "dom_find"
}

func (d *dom_find) Do(a configmodel.Action) string {
	t := getTempTarget(a)
	if t != nil {
		d, ok := t.(*parser.Dom)
		if ok {
			v := d.Find(a.Operation.Value).Text()
			setTempData(a.Operation.Key, v)

		}

	}
	return ""
}
