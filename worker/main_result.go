package worker

type Result struct {
}

func (r Result) StatusCode() int {

	//408 timeout 等网络传输错误, 返回-60 (60s)
	//50x, 返回-60 (60s)
	//403, -120 (120s)



}
