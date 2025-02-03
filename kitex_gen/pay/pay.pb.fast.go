// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package pay

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *PayReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PayReq[number], err)
}

func (x *PayReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TradeNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PayReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Subject, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PayReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.TotalAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PayReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.PayType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PayRsp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PayRsp[number], err)
}

func (x *PayRsp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PayRsp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PayRsp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.LogId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PayRsp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v PayRsp_PayRspData
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Data = &v
	return offset, nil
}

func (x *PayRsp_PayRspData) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PayRsp_PayRspData[number], err)
}

func (x *PayRsp_PayRspData) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PayPageUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PayRsp_PayRspData) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.TradeNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PayReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *PayReq) fastWriteField1(buf []byte) (offset int) {
	if x.TradeNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTradeNo())
	return offset
}

func (x *PayReq) fastWriteField2(buf []byte) (offset int) {
	if x.Subject == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSubject())
	return offset
}

func (x *PayReq) fastWriteField3(buf []byte) (offset int) {
	if x.TotalAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetTotalAmount())
	return offset
}

func (x *PayReq) fastWriteField4(buf []byte) (offset int) {
	if x.PayType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPayType())
	return offset
}

func (x *PayRsp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *PayRsp) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *PayRsp) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *PayRsp) fastWriteField3(buf []byte) (offset int) {
	if x.LogId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetLogId())
	return offset
}

func (x *PayRsp) fastWriteField4(buf []byte) (offset int) {
	if x.Data == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetData())
	return offset
}

func (x *PayRsp_PayRspData) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *PayRsp_PayRspData) fastWriteField1(buf []byte) (offset int) {
	if x.PayPageUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPayPageUrl())
	return offset
}

func (x *PayRsp_PayRspData) fastWriteField2(buf []byte) (offset int) {
	if x.TradeNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTradeNo())
	return offset
}

func (x *PayReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *PayReq) sizeField1() (n int) {
	if x.TradeNo == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTradeNo())
	return n
}

func (x *PayReq) sizeField2() (n int) {
	if x.Subject == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSubject())
	return n
}

func (x *PayReq) sizeField3() (n int) {
	if x.TotalAmount == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetTotalAmount())
	return n
}

func (x *PayReq) sizeField4() (n int) {
	if x.PayType == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetPayType())
	return n
}

func (x *PayRsp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *PayRsp) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetCode())
	return n
}

func (x *PayRsp) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *PayRsp) sizeField3() (n int) {
	if x.LogId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetLogId())
	return n
}

func (x *PayRsp) sizeField4() (n int) {
	if x.Data == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetData())
	return n
}

func (x *PayRsp_PayRspData) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *PayRsp_PayRspData) sizeField1() (n int) {
	if x.PayPageUrl == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPayPageUrl())
	return n
}

func (x *PayRsp_PayRspData) sizeField2() (n int) {
	if x.TradeNo == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTradeNo())
	return n
}

var fieldIDToName_PayReq = map[int32]string{
	1: "TradeNo",
	2: "Subject",
	3: "TotalAmount",
	4: "PayType",
}

var fieldIDToName_PayRsp = map[int32]string{
	1: "Code",
	2: "Msg",
	3: "LogId",
	4: "Data",
}

var fieldIDToName_PayRsp_PayRspData = map[int32]string{
	1: "PayPageUrl",
	2: "TradeNo",
}
