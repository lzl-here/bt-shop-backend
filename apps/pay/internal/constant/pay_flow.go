package constant

const (
	PayStaus           = iota
	PayStatusSuccess   = "SUCCESS"
	PayStatusCancel    = "CANCEL"
	PayStatusRefund    = "REFUND"
	PayStatusCompleted = "COMPLETED"
)

const (
	PayType = iota
	PayTypeAlipay = "ALIPAY"
)
