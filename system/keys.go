package system

var (
	AuthApiPath = "/auth"
	ListApiPath = "/queue"
	GetApiPath  = "/get"
	TplApiPath  = "/tpl"

	Method404Code = "NOT_FOUND"
	Method404Msg  = "Not found"

	ActionTempPoolName = "temp"
	SecondSleep        = 60 * 5
	MaxError           = "max_error"

	SystemGateway = "gateway"
	SystemPort    = "port"
	SystemApiAuth = "auth"
	SystemToken   = "token"
)
