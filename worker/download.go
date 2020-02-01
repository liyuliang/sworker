package worker

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/liyuliang/configmodel"
	"github.com/liyuliang/utils/regex"
	"github.com/liyuliang/utils/request"
	"bytes"
	"io/ioutil"
)

func init() {

	Register(func() worker {
		return new(download)
	})
}

type download struct {
}

func (d *download) Name() string {
	return "download"
}

func (d *download) Before(a configmodel.Action) configmodel.Action {
	v := a.Target.Value

	if a.Before.Replace.Target == "url" {
		v = regex.Replace(v, a.Before.Replace.From, a.Before.Replace.To)
	}
	a.Target.Value = v
	return a
}

func (d *download) Do(a configmodel.Action) {

	if a.Target.Key != "url" {
		return
	}

	a = d.Before(a)

	resp := request.HttpGet(a.Target.Value)

	statusCode = resp.StatusCode

	if resp.StatusCode != 200 {
		return
	}

	v := resp.Data
	if a.After.Transform.Target == "gbk" {
		v = gbkToUtf8(v)
	}

	if a.After.Replace.Target == "html" {
		v = regex.Replace(v, a.After.Replace.From, a.After.Replace.To)
	}
	ioutil.WriteFile("a.html", []byte(resp.Data), 0644)

	setTempData(a.Target.Key, a.Target.Value)
	setTempData(a.Operation.Key, v)
	return
}

func gbkToUtf8(text string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(text)), simplifiedchinese.GBK.NewDecoder())
	d, _ := ioutil.ReadAll(reader)
	return string(d)
}
