package worker

import (
	"github.com/liyuliang/configmodel"
	"net/url"
)

func init() {

	Register(func() worker {
		return new(url_host)
	})
}

type url_host struct {
}

func (d *url_host) Name() string {
	return "url_host"
}

func (d *url_host) Do(a configmodel.Action) {
	t := getTempTarget(a)
	if t != nil {
		s, ok := t.(string)
		if ok {
			u, err := url.Parse(s)
			if err == nil {
				host := u.Scheme + "://" + u.Host
				setTempData(a.Operation.Key, host)
			}
		}
	}
}
