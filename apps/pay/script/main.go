package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/transport"
	etcd "github.com/kitex-contrib/registry-etcd"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/example"
	pc "github.com/lzl-here/bt-shop-backend/kitex_gen/example/exampleservice"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}, etcd.WithAuthOpt("root", "root"))
	if err != nil {
		panic(err)
	}
	cli, err := pc.NewClient("example", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		// client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	if err != nil {
		panic(err)
	}
	var resp *pgen.GetPersonRsp
	resp, err = cli.GetPerson(context.TODO(), &pgen.GetPersonReq{Uid: 1})
	if bizErr, ok := kerrors.FromBizStatusError(err); ok {
		println(bizErr.BizStatusCode(), bizErr.BizMessage(), bizErr.BizExtra())
	}
	fmt.Println(resp)
}
