package worker

import (
	"github.com/liyuliang/configmodel"
	"sync"
	"strings"
)

type worker interface {
	Name() string
	Do(a configmodel.Action)
}

type Creator func() worker

var _list []Creator

func Register(method Creator) {
	_list = append(_list, method)
}

type parserList struct {
	sync.RWMutex
	creators map[string]Creator
}

var list parserList

func List() map[string]Creator {

	if len(list.creators) != len(_list) {

		list = parserList{}
		list.creators = make(map[string]Creator)
		list.Lock()

		for _, agent := range _list {
			list.creators[agent().Name()] = agent
		}
		list.Unlock()
	}

	return list.creators
}

func Get(name string) (creator Creator) {
	for _, agent := range List() {
		if strings.ToLower(agent().Name()) == strings.ToLower(name) {
			creator = agent
			break
		}
	}
	return creator
}

