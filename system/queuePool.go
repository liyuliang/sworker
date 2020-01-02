package system

import (
	"github.com/liyuliang/configmodel"
	"time"
	"strings"
	"sort"
)

type queues struct {
	pool pool
}

type queue struct {
	Name          string
	weight        int
	Online        bool
	errorCount    int
	errorLastTime time.Time
}

func (q *queue) SetWeight(w int) {
	q.weight = w
}
func (q *queue) Weight() int {
	return q.weight
}
func (q *queue) PullJobs() []configmodel.Action {
	return []configmodel.Action{}
}
func (q *queue) Downgrade(num int) {

	//		1.发现执行失败超过阀值(时间段内超过阀值)
	//		2.队列为空

	//now := time.Now()
	switch num {

	case -1:

		q.weight--
	case -10:
		q.weight = q.weight - 10
	}
}

type pool []*queue

func (p pool) Len() int {
	return len(p)
}
func (p pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pool) Less(i, j int) bool {
	return p[j].weight < p[i].weight
}

var qs *queues

func Queues() *queues {

	if qs == nil {
		qs = new(queues)
	}
	return qs
}

func (qs *queues) Pool() pool {

	var p []*queue
	for _, q := range qs.pool {
		if q.weight > 0 {
			p = append(p, q)
		}
	}

	sort.Sort(pool(p)) // 按照 weight 的逆序排序

	return p
}

func (qs *queues) Get(name string) (q *queue) {

	exist := false
	for _, q = range qs.Pool() {
		if strings.ToLower(q.Name) == strings.ToLower(name) {
			exist = true
			break
		}
	}

	if !exist {
		q = new(queue)
		q.Name = name
		q.weight = 1
		q.Online = true

		qs.pool = append(qs.pool, q)
	}
	return q
}
func (qs *queues) Count() int {

	return len(qs.Pool())
	//c := 0
	//
	//for _, q := range qs.Pool() {
	//	if q.weight > 0 {
	//		c++
	//	}
	//}
	//return c
}
