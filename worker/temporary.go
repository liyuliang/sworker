package worker

import "github.com/liyuliang/configmodel"

func init() {
	Register(func() worker {
		return new(tempoary)
	})

}

type tempoary struct {
}

func (d *tempoary) Name() string {
	return "tempoary"
}

func (d *tempoary) Do(a configmodel.Action) {
}
