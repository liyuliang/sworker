package system

var (
	QueuePrefix      = "queue:"
	QueueTotalPrefix = "total:"

	AuthApiPath   = "/auth"
	ListApiPath   = "/queue"
	AddApiPath    = "/add"
	SubmitApiPath = "/submit"
	GetApiPath    = "/get"
	TplApiPath    = "/tpl"

	CategoryInAddApi = "type"

	Method404Code = "NOT_FOUND"
	Method404Msg  = "Not found"

	ActionTempPoolName   = "temp"
	ActionReturnPoolName = "return"
	MaxError             = "max_error"
	SecondSleep          = "sleep"

	SystemGateway = "gateway"
	SystemPort    = "port"
	SystemApiAuth = "auth"
	SystemToken   = "token"
)

var (
	//EmptyQueueWait = 60 * 5
	EmptyQueueWait = 5
)
