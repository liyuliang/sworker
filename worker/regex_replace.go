package worker

import (
	"github.com/liyuliang/configmodel"
)

func init() {
	Register(func() worker {
		return new(regex_replace)
	})
}

type regex_replace struct {
}

func (d *regex_replace) Name() string {
	return "regex_replace"
}

func (d *regex_replace) Do(a configmodel.Action) {

	//TODO
	//getTempTarget()
}
