package system

import (
	"github.com/liyuliang/utils/format"
	"time"
	"strings"
	"sort"
	"github.com/liyuliang/utils/request"
	"encoding/json"
	"log"
)

type queues struct {
	pool pool
}

type queue struct {
	Name   string
	weight int
	//pendingWeight   int
	//errorCount      int
	//lastErrorTime   time.Time
	lastRestoreTime time.Time
}

func (q *queue) SetWeight(w int) {
	q.weight = w
}

func (q *queue) Weight() int {
	return q.weight
}

func (q *queue) PullTasks() (tasks []Task) {

	gateway := Config()[SystemGateway]
	queueGetApi := gateway + GetApiPath

	data := format.ToMap(map[string]string{
		"queue": q.Name,
		"n":     "1",
	})
	html, err := request.HttpPost(queueGetApi, data.ToUrlVals())

	if err != nil {
		return
	}
	log.Println("task response:")
	log.Println(queueGetApi, data.String())

	var urls []string
	json.Unmarshal([]byte(html), &urls)

	for _, u := range urls {
		log.Println(u)
		if u != "" {
			t := Task{
				Url: u,
			}
			tasks = append(tasks, t)
		}
	}

	return tasks
}

func (q *queue) Downgrade10min() {
	q.Downgrade(-60 * 10)
}

func (q *queue) Downgrade60min() {
	q.Downgrade(-60 * 60)
}

func (q *queue) Downgrade2hour() {
	q.Downgrade(-60 * 60 * 2)
}
func (q *queue) Downgrade24hour() {
	q.Downgrade(-60 * 60 * 24)
}

func (q *queue) Downgrade(num int) {

	//		1.发现执行失败超过阀值(时间段内超过阀值)
	//		2.队列为空

	if num < 0 {
		q.weight = q.weight + num
	}
}

func (q *queue) NaturalRestore() {

	now := time.Now()
	minute := 60
	expired := q.lastRestoreTime.Add(format.IntToTime(minute))

	if q.Weight() < 1 && expired.Before(now) {
		q.weight++
		q.lastRestoreTime = time.Now()
	}
}

func (q *queue) Online() bool {
	return q.weight > 0
}

func (q *queue) ResetWeight() {
	q.weight = 0
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

		qs.pool = append(qs.pool, q)
	}
	return q
}

func (qs *queues) Count() int {

	return len(qs.Pool())
}
