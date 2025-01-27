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

func (x *PrepayReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *PrepayRsp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PrepayRsp[number], err)
}

func (x *PrepayRsp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PrepayRsp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PrepayRsp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.LogId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PrepayRsp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v PrepayRsp_PrepayRspData
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Data = &v
	return offset, nil
}

func (x *PrepayRsp_PrepayRspData) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *PrepayReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *PrepayRsp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *PrepayRsp) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *PrepayRsp) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *PrepayRsp) fastWriteField3(buf []byte) (offset int) {
	if x.LogId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetLogId())
	return offset
}

func (x *PrepayRsp) fastWriteField4(buf []byte) (offset int) {
	if x.Data == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetData())
	return offset
}

func (x *PrepayRsp_PrepayRspData) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *PrepayReq) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *PrepayRsp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *PrepayRsp) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetCode())
	return n
}

func (x *PrepayRsp) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *PrepayRsp) sizeField3() (n int) {
	if x.LogId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetLogId())
	return n
}

func (x *PrepayRsp) sizeField4() (n int) {
	if x.Data == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetData())
	return n
}

func (x *PrepayRsp_PrepayRspData) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_PrepayReq = map[int32]string{}

var fieldIDToName_PrepayRsp = map[int32]string{
	1: "Code",
	2: "Msg",
	3: "LogId",
	4: "Data",
}

var fieldIDToName_PrepayRsp_PrepayRspData = map[int32]string{}
