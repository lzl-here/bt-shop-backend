package constant

// 待付款 --(完成付款)-->  已付款 --(部分订单项发货)--> 部分发货 --(全部订单项发货)--> 已发货 --(全部订单项签收)--> 已完成
// 待付款 --(取消付款)--> 已取消
const (
	OrderStatusPaying      = "paying"       // 待付款
	OrderStatusPaid        = "paid"         // 已付款
	OrderStatusPartDelivered = "part_delivered" // 部分发货
	OrderStatusDelivered     = "delivered"      // 已发货
	OrderStatusCompleted   = "completed"    // 已完成
	OrderStatusCanceled    = "canceled"     // 已取消
)

// 支付方式
const (
	OrderPayTypeAlipay = "ALIPAY"
	OrderPayTypeWechat = "WECHAT"
)
