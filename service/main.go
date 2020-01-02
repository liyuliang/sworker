package service

import (
	"queue-services"
	"github.com/liyuliang/sworker/system"
	"encoding/json"
	"github.com/liyuliang/utils/request"
	"github.com/liyuliang/sworker/worker"
	"time"
	"github.com/liyuliang/utils/format"
	"log"
)

func Start() {

	initQueue()

	services.AddSingleProcessTask("Pull Job", func(workerNum int) (err error) {
		pullFromQueue()
		return
	})
}

func pullFromQueue() {

	//gateway := system.Config()[system.SystemGateway]
	//queueGetApi := gateway + system.GetApiPath

	qs := system.Queues()

	if qs.Count() == 0 {

		//		1.发现执行失败超过阀值(时间段内超过阀值)
		//		2.队列为空

		pending()

	} else {

		for _, q := range qs.Pool() {

			jobs := q.PullJobs()

			if len(jobs) == 0 {

				q.Downgrade(-1)
				continue
			}

			//jobs := pullJobs(q.Name)

			for _, job := range jobs {

				result := worker.Run(job)

				if result.Code() < 1 {
					q.Downgrade(result.Code())
				}
			}
		}
	}
}

func pending() {
	sleep := format.StrToInt(system.SecondSleep)
	log.Printf("Empty queue, wait %d second", sleep)
	time.Sleep(format.IntToTimeSecond(sleep))
}

func downgrade(queue string, r worker.Result) {

	//
	switch r.Code() {

	case 0:

		queue.offline()
	default:
		queue.downgrade(r.Code())
	}
}

func initQueue() {

	gateway := system.Config()[system.SystemGateway]
	queueListApi := gateway + system.ListApiPath

	resp := request.HttpGet(queueListApi)
	data := make(map[string]string)

	json.Unmarshal([]byte(resp.Data), &data)

	q := system.Queues()

	//{
	//    "queue_list_gufengmh8": "12"
	//}

	for queueName, count := range data {

		q.Get(queueName).SetWeight(format.StrToInt(count))
	}

}
