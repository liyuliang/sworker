package main

import (
	"github.com/liyuliang/sworker/route"
	"flag"
)

func main() {

	//cli 中获取 -a, --auth 提交到队列中心校验身份
	//cli 中获取 -port, 开启 web 界面(获取 top 信息)
	//程序启动
	//上报 ip、top、启动时间
	//请求爬虫任务队列
	//根据任务队列优先级{队列名称:优先级}, 获取任务
	//if empty { next queue }
	//if current_queue_max_failed { next queue }
	//if no_available_queue { hold on }

	route.Start(p)
}

var (
	a string
	p string
)

func init() {
	flag.StringVar(&a, "a", "", "auth token")
	flag.StringVar(&p, "p", "8989", "web port")

	flag.Parse()
}
