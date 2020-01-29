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
	"github.com/liyuliang/queue-services"
	"strings"
	"math"
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

	services.AddSingleProcessTask("Report self profile", func(workerNum int) (err error) {
		//TODO report profile to update token
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

			var taskResult []worker.Data
			var extendResult []worker.Data

			for _, t := range tasks {

				actions := genActions(q.Name, t)

				if len(actions) == 0 {
					continue
				}

				worker.Clean()

				for _, a := range actions {
					worker.Run(a)
				}

				q.ChangeWeightByStatusCode(worker.StatusCode())


				result := worker.ReturnData()
				taskResult = append(taskResult, result)


				extend := worker.NewData()
				extendResult = append(extendResult, extend)
			}

			if len(taskResult) > 0 {
				Submit(taskResult)
			}

			if len(extendResult) > 0 {
				Submit(extendResult)
			}
		}
	}
}

func Submit(data []worker.Data) {

	gateway := system.Config()[system.SystemGateway]
	queueSubmitApi := gateway + system.SubmitApiPath

	chunk := sliceChunk(data, 10)

	for _, slice := range chunk {

		for _, actions := range slice {

			request.HttpPost(queueSubmitApi, actions.ToUrlVals())
		}
	}
}

func Extend(data []worker.Data) {
	gateway := system.Config()[system.SystemGateway]

	queueAddApi := gateway + system.AddApiPath

	chunk := sliceChunk(data, 10)

	for _, slice := range chunk {

		for _, actions := range slice {

			request.HttpPost(queueAddApi, actions.ToUrlVals())
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

	for queueName, count := range data {

		q.Get(queueName).SetWeight(format.StrToInt(count))
	}
}

func sliceChunk(data []worker.Data, size int) (result [][]worker.Data) {

	l := len(data)
	if l < size {
		result = append(result, data)
		return result
	}

	groupLen := int(math.Ceil(float64(l/size))) + 1

	for i := 0; i < groupLen; i++ {

		start := i * size
		end := start + size
		var newSlice []worker.Data

		log.Println(len(data))
		log.Println(start)
		log.Println(end)
		newSlice = data[start:end]

		result = append(result, newSlice)
	}
	return result
}
