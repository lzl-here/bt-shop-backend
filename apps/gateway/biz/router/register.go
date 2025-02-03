// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api_pay "github.com/lzl-here/bt-shop-backend/apps/gateway/biz/router/api_pay"
	user "github.com/lzl-here/bt-shop-backend/apps/gateway/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	user.Register(r)

	api_pay.Register(r)

}
