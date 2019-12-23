package system

import (
	"github.com/liyuliang/utils/request"
	"strings"
	"os"
	"fmt"
)

type appConfig map[string]string

var _config appConfig

func init() {
	_config = make(map[string]string)
}

func Init(gateway, auth string) {

	//gateway
	api := "http://localhost:7777/api/auth?key=" + auth
	resp := request.HttpGet(api)
	if !strings.Contains(resp.Data, "success") {
		fmt.Fprintf(os.Stderr, "Auth failed \n")
		os.Exit(2)
	}






}
