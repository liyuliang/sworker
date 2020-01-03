package worker

import (
	"github.com/liyuliang/configmodel"
	"github.com/liyuliang/dom-parser"
)

func init() {

	Register(func() worker {
		return new(dom_find_all)
	})
}

type dom_find_all struct {
}

func (d *dom_find_all) Name() string {
	return "dom_find_all"
}

func (d *dom_find_all) Do(a configmodel.Action)  {

	var data []string

	t := getTempTarget(a)
	if t != nil {
		dom, ok := t.(*parser.Dom)
		if ok {
			for _, d := range dom.FindAll(a.Operation.Value) {

				if a.Operation.Option.Key == "" {
					v := d.Find(a.Operation.Value).Text()
					data = append(data, v)
				} else {

					switch a.Operation.Option.Type {
					case "attr":
						v, exist := d.Attr(a.Operation.Option.Key)
						if exist {
							data = append(data, v)
						}
					}
				}
			}
			setTempData(a.Operation.Key, data)
		}
	}
	return
}
