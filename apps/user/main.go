package main

import (
	"github.com/bt-shop-backend/apps/user/biz/Dail/mysql"
	userSrv "github.com/lzl-here/bt-shop-backend/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := userSrv.NewServer(new(UserServiceImpl))
	mysql.Init()
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
