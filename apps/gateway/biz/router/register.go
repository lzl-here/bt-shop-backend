// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	goods "github.com/lzl-here/bt-shop-backend/apps/gateway/biz/router/goods"
	order "github.com/lzl-here/bt-shop-backend/apps/gateway/biz/router/order"
	pay "github.com/lzl-here/bt-shop-backend/apps/gateway/biz/router/pay"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	pay.Register(r)

	goods.Register(r)

	order.Register(r)

}
