package service

import (
	"github.com/liyuliang/sworker/system"
	"github.com/liyuliang/sworker/worker"
	"github.com/liyuliang/utils/request"
	"github.com/liyuliang/utils/format"
	"encoding/json"
	"time"
	"log"
	"github.com/liyuliang/configmodel"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/liyuliang/queue-services"
	"strings"
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
	services.Service().Start(false)
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

		for _, q := range qs.Pool() {

			tasks := q.PullTasks()

			if len(tasks) == 0 {

				q.ResetWeight()
				q.Downgrade10min()
				continue
			}

			stopJobs := false
			for _, t := range tasks {

				actions := genActions(q.Name, t)

				log.Println("Action count:", len(actions))
				worker.Clean()

				for _, a := range actions {
					worker.Run(a)
				}

				data := worker.ReturnData()

				fmt.Print(data)
				switch worker.StatusCode() {

				case 0:
				case 200:

					//Do nothing
				case 403:
					q.ResetWeight()
					q.Downgrade60min()

					stopJobs = true
				default:
					q.Downgrade10min()
				}

				if stopJobs {
					break
				}
			}
		}
	}
}
func genActions(queueName string, task system.Task) (as []configmodel.Action) {
	queueName = strings.Replace(queueName, system.QueuePrefix, "", -1)

	tpl := system.Config()[queueName]
	model := new(configmodel.Actions)
	_, err := toml.Decode(tpl, &model)

	if err != nil {
		log.Println(err.Error())
		return
	}

	for i, a := range model.Action {
		if a.Target.Key == "ur" && a.Operation.Type == "download" {
			model.Action[i].Target.Value = task.Url
		}
	}
	as = model.Action
	return as
}

func pending() {
	sleep := system.EmptyQueueWait
	log.Printf("Empty queue, wait %d second", sleep)
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
