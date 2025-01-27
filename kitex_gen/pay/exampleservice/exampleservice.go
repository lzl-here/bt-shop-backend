// Code generated by Kitex v0.12.1. DO NOT EDIT.

package exampleservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	pay "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Prepay": kitex.NewMethodInfo(
		prepayHandler,
		newPrepayArgs,
		newPrepayResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	exampleServiceServiceInfo                = NewServiceInfo()
	exampleServiceServiceInfoForClient       = NewServiceInfoForClient()
	exampleServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return exampleServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return exampleServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return exampleServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ExampleService"
	handlerType := (*pay.ExampleService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "pay",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func prepayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(pay.PrepayReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(pay.ExampleService).Prepay(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PrepayArgs:
		success, err := handler.(pay.ExampleService).Prepay(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PrepayResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPrepayArgs() interface{} {
	return &PrepayArgs{}
}

func newPrepayResult() interface{} {
	return &PrepayResult{}
}

type PrepayArgs struct {
	Req *pay.PrepayReq
}

func (p *PrepayArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(pay.PrepayReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PrepayArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PrepayArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PrepayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PrepayArgs) Unmarshal(in []byte) error {
	msg := new(pay.PrepayReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PrepayArgs_Req_DEFAULT *pay.PrepayReq

func (p *PrepayArgs) GetReq() *pay.PrepayReq {
	if !p.IsSetReq() {
		return PrepayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PrepayArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PrepayArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PrepayResult struct {
	Success *pay.PrepayRsp
}

var PrepayResult_Success_DEFAULT *pay.PrepayRsp

func (p *PrepayResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(pay.PrepayRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PrepayResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PrepayResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PrepayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PrepayResult) Unmarshal(in []byte) error {
	msg := new(pay.PrepayRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PrepayResult) GetSuccess() *pay.PrepayRsp {
	if !p.IsSetSuccess() {
		return PrepayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PrepayResult) SetSuccess(x interface{}) {
	p.Success = x.(*pay.PrepayRsp)
}

func (p *PrepayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrepayResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Prepay(ctx context.Context, Req *pay.PrepayReq) (r *pay.PrepayRsp, err error) {
	var _args PrepayArgs
	_args.Req = Req
	var _result PrepayResult
	if err = p.c.Call(ctx, "Prepay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
