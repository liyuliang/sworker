package worker

import "github.com/liyuliang/configmodel"

func init() {
	Register(func() worker {
		return new(temporary)
	})
}

type temporary struct {
}

func (d *temporary) Name() string {
	return "temporary"
}

func (d *temporary) Do(a configmodel.Action) {

	setTempData(a.Operation.Key, a.Operation.Value)
}
