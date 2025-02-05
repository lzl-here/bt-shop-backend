package common

import (
	"strconv"

	"github.com/smartwalle/xid"
)

func genID() int64 {
	return xid.Next()
}
func GenTradeNo() string {
	id := genID()
	return "t_" + strconv.FormatInt(id, 10)
}

func GenOrderNo() string {
	id := genID()
	return "o_" + strconv.FormatInt(id, 10)
}
