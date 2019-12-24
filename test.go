package main

import (
	"github.con/liyuliang/configmodel"
	"github.com/BurntSushi/toml"
	"os"
)

func main() {

	as := new(configmodel.Actions)
	_, err := toml.DecodeFile("./test.toml", as)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	for _, action := range as.Action {
		println(action.Target.Key)
	}
}
