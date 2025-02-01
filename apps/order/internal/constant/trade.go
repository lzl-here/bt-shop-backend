package constant

// 待付款 --(完成付款)--> 已付款 --(所有订单完成发货)--> 已完成
// 待付款 --(取消付款)--> 已取消
const (
	TradeStatusPaying    = "paying"    // 待付款
	TradeStatusPaid      = "paid"      // 已付款
	TradeStatusCompleted = "completed" // 已完成
	TradeStatusCancel    = "cancel"    // 已取消
)
