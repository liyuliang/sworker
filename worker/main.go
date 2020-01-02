package worker

import (
	"github.com/liyuliang/configmodel"
	"log"
)

func Run(a configmodel.Action) Result {

	if a.Target.Value != "" {
		setTempData(a.Target.Key, a.Target.Value)
	}

	w := Get(a.Operation.Type)
	if w == nil {
		log.Printf("worker %s is missing from operation type", a.Operation.Type)
		return
	}
	log.Printf("worker %s prepare run", w().Name())
	w().Do(a)

	data := getTempData(a.Return)
	if a.Return != "" && data != "" {
		replaceReturnData(a.Return, data)
	}

}
