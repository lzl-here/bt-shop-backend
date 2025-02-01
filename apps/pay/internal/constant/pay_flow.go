package constant

const (
	PayStaus           = iota
	PayStatusSuccess   = "SUCCESS"
	PayStatusCancel    = "CANCEL"
	PayStatusRefund    = "REFUND"
	PayStatusCompleted = "COMPLETED"
)

const (
	PayType       = iota
	PayTypeAlipay = "ALIPAY"
)

const (
	PayInType  = "in"  // 入帐
	PayOutType = "out" // 出账
)
