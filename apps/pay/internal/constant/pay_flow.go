package constant

const (
	PayStatePaying       = 1 // 支付中
	PayStatePayed        = 2 // 已支付
	PayStateCompleted    = 3 // 完成结算
	PayStateCanceled     = 4 // 取消支付
	PayStateRefunding    = 5 // 退款中
	PayStateFullRefunded = 6 // 全额退款
	PayStatePartRefunded = 7 // 部分退款
	PayStateRefundFailed = 8 // 退款失败
)
