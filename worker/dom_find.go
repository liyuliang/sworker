package worker

import (
	"github.com/liyuliang/configmodel"
	"github.com/liyuliang/dom-parser"
	"strings"
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

func (d *dom_find) Do(a configmodel.Action) {
	t := getTempTarget(a)
	if t != nil {
		d, ok := t.(*parser.Dom)
		if ok {

			if a.Operation.Option.Key == "" {
				v := d.Find(a.Operation.Value).Text()
				v = removeLineBreak(v)
				setTempData(a.Operation.Key, v)

			} else {

				switch a.Operation.Option.Type {
				case "attr":
					v, _ := d.Find(a.Operation.Value).Attr(a.Operation.Option.Key)
					setTempData(a.Operation.Key, v)
				}
			}

		}
	}
	return
}

func removeLineBreak(v string) string {
	v = strings.Replace(v, "\n", "", -1)
	v = strings.Replace(v, "\t", "", -1)
	v = strings.Replace(v, " ", "", -1)
	return v
}
