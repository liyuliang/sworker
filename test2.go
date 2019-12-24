package main
//
//import (
//	"github.com/BurntSushi/toml"
//	"os"
//	"github.con/liyuliang/configmodel"
//	//"configmodel"
//	"bytes"
//	"io/ioutil"
//	"golang.org/x/text/transform"
//	"golang.org/x/text/encoding/simplifiedchinese"
//	"fmt"
//	"strings"
//	"log"
//	"github.com/astaxie/beego/logs"
//	"github.com/liyuliang/utils/regex"
//	"github.com/liyuliang/utils/request"
//	"github.com/liyuliang/dom-parser"
//)
//
//var tempData map[string]interface{}
//var returnData map[string]interface{}
//
//func main() {
//
//	as := new(configmodel.Actions)
//	_, err := toml.DecodeFile("./test.toml", as)
//	if err != nil {
//		println(err.Error())
//		os.Exit(-1)
//	}
//
//
//	tempData = make(map[string]interface{})
//	returnData = make(map[string]interface{})
//
//
//	for _, action := range as.Action {
//
//
//		if action.Target.Value != "" {
//			tempData[action.Target.Key] = action.Target.Value
//		}
//
//		run(action)
//
//
//		if action.Return != "" {
//			returnData[action.Return] = tempData[action.Return]
//		}
//	}
//	fmt.Print(returnData)
//
//}
//func run(action configmodel.Action) {
//	switch strings.ToLower(action.Operation.Type) {
//
//	case "download":
//
//		if action.Target.Key == "url" {
//
//			v := action.Target.Value
//
//			if action.Before.Replace.Target == "url" {
//				v = regex.Replace(v, action.Before.Replace.From, action.Before.Replace.To)
//			}
//
//			log.Println(v)
//			resp := request.HttpGet(v)
//
//			respV := resp.Data
//
//			err := ioutil.WriteFile("a.html", []byte(respV), 0644)
//			if err != nil {
//				println(err.Error())
//			}
//			if action.After.Transform.Target == "gbk" {
//				respV = gbkToUtf8(respV)
//			}
//
//			if action.After.Replace.Target == "html" {
//				respV = regex.Replace(respV, action.After.Replace.From, action.After.Replace.To)
//
//			}
//
//			tempData[action.Operation.Key] = respV
//		}
//	case "regex\\.get":
//
//		v, ok := tempData[action.Target.Key].(string)
//		if ok {
//			v = regex.Get(v, action.Operation.Value)
//			tempData[action.Target.Key] = v
//			tempData[action.Operation.Key] = v
//		}
//	case "dom_parser":
//
//		if action.Target.Type == "temp" {
//			v, ok := tempData[action.Target.Key].(string)
//			if ok {
//				d, e := parser.InitDom(v)
//				if e != nil {
//
//					logs.Error(e.Error())
//				} else {
//					tempData[action.Operation.Key] = d
//				}
//			}
//		}
//	case "find":
//		println(action.Target.Type )
//		println(action.Target.Key)
//
//		if action.Target.Type == "temp" {
//			d, ok := tempData[action.Target.Key].(parser.Dom)
//			if ok {
//				println("ok")
//				v := d.Find(action.Operation.Value).Text()
//				log.Println(v)
//				tempData[action.Operation.Key] = v
//
//			}
//
//		}
//	}
//
//}
//
//func gbkToUtf8(text string) string {
//	reader := transform.NewReader(bytes.NewReader([]byte(text)), simplifiedchinese.GBK.NewDecoder())
//	d, _ := ioutil.ReadAll(reader)
//	return string(d)
//}
