// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package order

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *CancelTradeReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CancelTradeReq[number], err)
}

func (x *CancelTradeReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TradeNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CancelTradeRsp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CancelTradeRsp[number], err)
}

func (x *CancelTradeRsp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CancelTradeRsp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CancelTradeRsp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.LogId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CancelTradeRsp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v CancelTradeRsp_CancelTradeRspData
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Data = &v
	return offset, nil
}

func (x *CancelTradeRsp_CancelTradeRspData) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CancelTradeRsp_CancelTradeRspData[number], err)
}

func (x *CancelTradeRsp_CancelTradeRspData) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TradeNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CancelTradeReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CancelTradeReq) fastWriteField1(buf []byte) (offset int) {
	if x.TradeNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTradeNo())
	return offset
}

func (x *CancelTradeRsp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *CancelTradeRsp) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *CancelTradeRsp) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *CancelTradeRsp) fastWriteField3(buf []byte) (offset int) {
	if x.LogId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetLogId())
	return offset
}

func (x *CancelTradeRsp) fastWriteField4(buf []byte) (offset int) {
	if x.Data == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetData())
	return offset
}

func (x *CancelTradeRsp_CancelTradeRspData) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CancelTradeRsp_CancelTradeRspData) fastWriteField1(buf []byte) (offset int) {
	if x.TradeNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTradeNo())
	return offset
}

func (x *CancelTradeReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CancelTradeReq) sizeField1() (n int) {
	if x.TradeNo == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTradeNo())
	return n
}

func (x *CancelTradeRsp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *CancelTradeRsp) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetCode())
	return n
}

func (x *CancelTradeRsp) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *CancelTradeRsp) sizeField3() (n int) {
	if x.LogId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetLogId())
	return n
}

func (x *CancelTradeRsp) sizeField4() (n int) {
	if x.Data == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetData())
	return n
}

func (x *CancelTradeRsp_CancelTradeRspData) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CancelTradeRsp_CancelTradeRspData) sizeField1() (n int) {
	if x.TradeNo == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTradeNo())
	return n
}

var fieldIDToName_CancelTradeReq = map[int32]string{
	1: "TradeNo",
}

var fieldIDToName_CancelTradeRsp = map[int32]string{
	1: "Code",
	2: "Msg",
	3: "LogId",
	4: "Data",
}

var fieldIDToName_CancelTradeRsp_CancelTradeRspData = map[int32]string{
	1: "TradeNo",
}
