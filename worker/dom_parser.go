package worker

import (
	"github.com/liyuliang/configmodel"
	"github.com/liyuliang/dom-parser"
	"log"
)

func init() {

	Register(func() worker {
		return new(dom_parser)
	})
}

type dom_parser struct {
}

func (d *dom_parser) Name() string {
	return "dom_parser"
}

func (d *dom_parser) Do(a configmodel.Action) {
	t := getTempTarget(a)
	if t != nil {
		v, ok := t.(string)
		if ok {
			d, e := parser.InitDom(v)
			if e != nil {

				log.Println(e.Error())
			} else {
				setTempData(a.Operation.Key, d)
			}
		}
	}
}
