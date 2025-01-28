package constant

const (
	PayStatePaying       = 1 // 支付中
	PayStatePayed        = 2 // 已支付
	PayStateCompleted    = 3 // 完成结算
	PayStateCanceled     = 4 // 取消取消
	PayStateRefunded     = 5 // 退款
	PayStatePartRefunded = 6 // 部分退款
	PayStateRefundFailed = 7 // 退款失败
)

