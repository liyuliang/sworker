package worker

type Result struct {
}

func (r Result) Code() int {

	//408 timeout 等网络传输错误, 返回1
	//50x, 返回-1
	//403, -10



}
