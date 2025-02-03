// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package goods

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *StockReduceReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *StockReduceRsp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *StockReduceRollbackReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *StockReduceRollbackRsp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *StockReduceReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *StockReduceRsp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *StockReduceRollbackReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *StockReduceRollbackRsp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *StockReduceReq) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *StockReduceRsp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *StockReduceRollbackReq) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *StockReduceRollbackRsp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_StockReduceReq = map[int32]string{}

var fieldIDToName_StockReduceRsp = map[int32]string{}

var fieldIDToName_StockReduceRollbackReq = map[int32]string{}

var fieldIDToName_StockReduceRollbackRsp = map[int32]string{}
