package constant

var (
	// AlipayTradeFinished = "TRADE_FINISHED" // 交易成功，商家不开启退款功能 (我们用不到，使用success这个就行)
	AlipayTradeSuccess  = "TRADE_SUCCESS"  // 交易成功，商家开启退款功能
	AlipayTradeClosed   = "TRADE_CLOSED"   // 未付款交易超时关闭，或支付完成后全额退款。
	// AlipayTradeWait     = "WAIT_BUYER_PAY" // 交易创建，等待买家付款，不采用
)


var (
	AlipayNoticeSyncTrade = "trade_status_sync"
)