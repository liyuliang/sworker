package service

import (
	"github.com/liyuliang/sworker/system"
	"github.com/liyuliang/sworker/worker"
	"github.com/liyuliang/utils/request"
	"github.com/liyuliang/utils/format"
	"encoding/json"
	"queue-services"
	"time"
	"log"
)

func Start() {

	initQueue()

	services.AddSingleProcessTask("Pull Job", func(workerNum int) (err error) {
		pullFromQueue()
		return
	})

	services.AddSingleProcessTask("Restore Queue Weight", func(workerNum int) (err error) {
		restoreQueueWeight()
		return
	})
}

func restoreQueueWeight() {
	qs := system.Queues()
	for _, q := range qs.Pool() {
		if q.Weight() < 1 {
			q.NaturalRestore()
		}
	}
}

func pullFromQueue() {

	qs := system.Queues()

	if qs.Count() == 0 {

		//全部队列为空
		pending()

	} else {
		//		1.发现执行失败超过阀值(时间段内超过阀值)
		//		2.队列为空

		for _, q := range qs.Pool() {

			jobs := q.PullJobs()

			if len(jobs) == 0 {

				q.ResetWeight()
				q.Downgrade10min()
				continue
			}

			for _, job := range jobs {

				resp := worker.Run(job)

				switch resp.StatusCode() {

				case 200:

					//Do nothing
				case 403:
					q.ResetWeight()
					q.Downgrade60min()

				default:
					q.Downgrade10min()
				}

			}
		}
	}
}

func pending() {
	sleep := system.SecondSleep
	log.Printf("Empty queue, wait %d second", system.SecondSleep)
	time.Sleep(format.IntToTimeSecond(sleep))
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
