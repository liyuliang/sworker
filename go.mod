module github.com/liyuliang/sworker

go 1.13

replace golang.org/x/sys v0.0.0-20190410235845-0ad05ae3009d => github.com/golang/sys v0.0.0-20190410235845-0ad05ae3009d

require (
	github.com/astaxie/beego v1.12.0
	github.com/gin-gonic/gin v1.5.0
	github.com/liyuliang/configmodel v0.0.0-20191224080310-8784ddec76cd
	github.com/liyuliang/dom-parser v0.0.0-20171020101219-83d71ce6d1ec
	github.com/liyuliang/utils v0.0.0-20190805150857-cdeb9c4f8ad0
	github.com/mackerelio/go-osstat v0.1.0
	golang.org/x/text v0.3.2
)
