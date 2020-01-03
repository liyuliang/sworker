package system

var (
	QueuePrefix      = "queue_"
	QueueTotalPrefix = "total_"

	AuthApiPath = "/auth"
	ListApiPath = "/queue"
	GetApiPath  = "/get"
	TplApiPath  = "/tpl"

	Method404Code = "NOT_FOUND"
	Method404Msg  = "Not found"

	ActionTempPoolName = "temp"
	MaxError           = "max_error"
	SecondSleep        = "sleep"

	SystemGateway = "gateway"
	SystemPort    = "port"
	SystemApiAuth = "auth"
	SystemToken   = "token"
)

var (
	EmptyQueueWait = 60 * 5
)
