package dal

import "github.com/lzl-here/bt-shop-backend/apps/user_gjw/biz/dal/mysql"

func Init() {
	// redis.Init()
	mysql.Init()
}
