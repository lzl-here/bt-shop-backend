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

func (x *GetGoodsListReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetGoodsListReq[number], err)
}

func (x *GetGoodsListReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint64
			v, offset, err = fastpb.ReadUint64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.SpuIds = append(x.SpuIds, v)
			return offset, err
		})
	return offset, err
}

func (x *GetGoodsListRsp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetGoodsListRsp[number], err)
}

func (x *GetGoodsListRsp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.LogId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v GetGoodsListRsp_GetGoodsListRspData
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Data = &v
	return offset, nil
}

func (x *GetGoodsListRsp_GetGoodsListRspData) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetGoodsListRsp_GetGoodsListRspData[number], err)
}

func (x *GetGoodsListRsp_GetGoodsListRspData) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v GetGoodsListRsp_SpuInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SpuList = append(x.SpuList, &v)
	return offset, nil
}

func (x *GetGoodsListRsp_SpuInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 12:
		offset, err = x.fastReadField12(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 13:
		offset, err = x.fastReadField13(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetGoodsListRsp_SpuInfo[number], err)
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SpuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpuName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SpuDesc, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SpuImgUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.SpuMinAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.SpuMaxAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.SpuCategoryId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.SpuCategoryName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.BrandId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	x.BrandName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	var v GetGoodsListRsp_SkuInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SkuList = append(x.SkuList, &v)
	return offset, nil
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField12(buf []byte, _type int8) (offset int, err error) {
	var v GetGoodsListRsp_SpecInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SpecList = append(x.SpecList, &v)
	return offset, nil
}

func (x *GetGoodsListRsp_SpuInfo) fastReadField13(buf []byte, _type int8) (offset int, err error) {
	var v GetGoodsListRsp_SpecValueInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SpecValueList = append(x.SpecValueList, &v)
	return offset, nil
}

func (x *GetGoodsListRsp_SkuInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetGoodsListRsp_SkuInfo[number], err)
}

func (x *GetGoodsListRsp_SkuInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SkuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SkuInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SkuInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.StockNum, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SkuInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SkuAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SkuInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.SkuName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SkuInfo) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.SpecStr, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpecInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetGoodsListRsp_SpecInfo[number], err)
}

func (x *GetGoodsListRsp_SpecInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SpecId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpecInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpecName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpecInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SpuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpecValueInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetGoodsListRsp_SpecValueInfo[number], err)
}

func (x *GetGoodsListRsp_SpecValueInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SpecValueId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpecValueInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpecName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListRsp_SpecValueInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SpecValue, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetGoodsListReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetGoodsListReq) fastWriteField1(buf []byte) (offset int) {
	if len(x.SpuIds) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.GetSpuIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint64(buf[offset:], numTagOrKey, x.GetSpuIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *GetGoodsListRsp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *GetGoodsListRsp) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *GetGoodsListRsp) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *GetGoodsListRsp) fastWriteField3(buf []byte) (offset int) {
	if x.LogId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetLogId())
	return offset
}

func (x *GetGoodsListRsp) fastWriteField4(buf []byte) (offset int) {
	if x.Data == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetData())
	return offset
}

func (x *GetGoodsListRsp_GetGoodsListRspData) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetGoodsListRsp_GetGoodsListRspData) fastWriteField1(buf []byte) (offset int) {
	if x.SpuList == nil {
		return offset
	}
	for i := range x.GetSpuList() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetSpuList()[i])
	}
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	offset += x.fastWriteField12(buf[offset:])
	offset += x.fastWriteField13(buf[offset:])
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField1(buf []byte) (offset int) {
	if x.SpuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSpuId())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField2(buf []byte) (offset int) {
	if x.SpuName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSpuName())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField3(buf []byte) (offset int) {
	if x.SpuDesc == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSpuDesc())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField4(buf []byte) (offset int) {
	if x.SpuImgUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSpuImgUrl())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField5(buf []byte) (offset int) {
	if x.SpuMinAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetSpuMinAmount())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField6(buf []byte) (offset int) {
	if x.SpuMaxAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetSpuMaxAmount())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField7(buf []byte) (offset int) {
	if x.SpuCategoryId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 7, x.GetSpuCategoryId())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField8(buf []byte) (offset int) {
	if x.SpuCategoryName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetSpuCategoryName())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField9(buf []byte) (offset int) {
	if x.BrandId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 9, x.GetBrandId())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField10(buf []byte) (offset int) {
	if x.BrandName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 10, x.GetBrandName())
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField11(buf []byte) (offset int) {
	if x.SkuList == nil {
		return offset
	}
	for i := range x.GetSkuList() {
		offset += fastpb.WriteMessage(buf[offset:], 11, x.GetSkuList()[i])
	}
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField12(buf []byte) (offset int) {
	if x.SpecList == nil {
		return offset
	}
	for i := range x.GetSpecList() {
		offset += fastpb.WriteMessage(buf[offset:], 12, x.GetSpecList()[i])
	}
	return offset
}

func (x *GetGoodsListRsp_SpuInfo) fastWriteField13(buf []byte) (offset int) {
	if x.SpecValueList == nil {
		return offset
	}
	for i := range x.GetSpecValueList() {
		offset += fastpb.WriteMessage(buf[offset:], 13, x.GetSpecValueList()[i])
	}
	return offset
}

func (x *GetGoodsListRsp_SkuInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *GetGoodsListRsp_SkuInfo) fastWriteField1(buf []byte) (offset int) {
	if x.SkuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSkuId())
	return offset
}

func (x *GetGoodsListRsp_SkuInfo) fastWriteField2(buf []byte) (offset int) {
	if x.SpuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetSpuId())
	return offset
}

func (x *GetGoodsListRsp_SkuInfo) fastWriteField3(buf []byte) (offset int) {
	if x.StockNum == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 3, x.GetStockNum())
	return offset
}

func (x *GetGoodsListRsp_SkuInfo) fastWriteField4(buf []byte) (offset int) {
	if x.SkuAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSkuAmount())
	return offset
}

func (x *GetGoodsListRsp_SkuInfo) fastWriteField5(buf []byte) (offset int) {
	if x.SkuName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetSkuName())
	return offset
}

func (x *GetGoodsListRsp_SkuInfo) fastWriteField6(buf []byte) (offset int) {
	if x.SpecStr == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetSpecStr())
	return offset
}

func (x *GetGoodsListRsp_SpecInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetGoodsListRsp_SpecInfo) fastWriteField1(buf []byte) (offset int) {
	if x.SpecId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSpecId())
	return offset
}

func (x *GetGoodsListRsp_SpecInfo) fastWriteField2(buf []byte) (offset int) {
	if x.SpecName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSpecName())
	return offset
}

func (x *GetGoodsListRsp_SpecInfo) fastWriteField3(buf []byte) (offset int) {
	if x.SpuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 3, x.GetSpuId())
	return offset
}

func (x *GetGoodsListRsp_SpecValueInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetGoodsListRsp_SpecValueInfo) fastWriteField1(buf []byte) (offset int) {
	if x.SpecValueId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSpecValueId())
	return offset
}

func (x *GetGoodsListRsp_SpecValueInfo) fastWriteField2(buf []byte) (offset int) {
	if x.SpecName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSpecName())
	return offset
}

func (x *GetGoodsListRsp_SpecValueInfo) fastWriteField3(buf []byte) (offset int) {
	if x.SpecValue == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSpecValue())
	return offset
}

func (x *GetGoodsListReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetGoodsListReq) sizeField1() (n int) {
	if len(x.SpuIds) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.GetSpuIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint64(numTagOrKey, x.GetSpuIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *GetGoodsListRsp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *GetGoodsListRsp) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetCode())
	return n
}

func (x *GetGoodsListRsp) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *GetGoodsListRsp) sizeField3() (n int) {
	if x.LogId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetLogId())
	return n
}

func (x *GetGoodsListRsp) sizeField4() (n int) {
	if x.Data == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetData())
	return n
}

func (x *GetGoodsListRsp_GetGoodsListRspData) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetGoodsListRsp_GetGoodsListRspData) sizeField1() (n int) {
	if x.SpuList == nil {
		return n
	}
	for i := range x.GetSpuList() {
		n += fastpb.SizeMessage(1, x.GetSpuList()[i])
	}
	return n
}

func (x *GetGoodsListRsp_SpuInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	n += x.sizeField12()
	n += x.sizeField13()
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField1() (n int) {
	if x.SpuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSpuId())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField2() (n int) {
	if x.SpuName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSpuName())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField3() (n int) {
	if x.SpuDesc == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSpuDesc())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField4() (n int) {
	if x.SpuImgUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSpuImgUrl())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField5() (n int) {
	if x.SpuMinAmount == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetSpuMinAmount())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField6() (n int) {
	if x.SpuMaxAmount == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetSpuMaxAmount())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField7() (n int) {
	if x.SpuCategoryId == 0 {
		return n
	}
	n += fastpb.SizeUint64(7, x.GetSpuCategoryId())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField8() (n int) {
	if x.SpuCategoryName == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetSpuCategoryName())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField9() (n int) {
	if x.BrandId == 0 {
		return n
	}
	n += fastpb.SizeUint64(9, x.GetBrandId())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField10() (n int) {
	if x.BrandName == "" {
		return n
	}
	n += fastpb.SizeString(10, x.GetBrandName())
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField11() (n int) {
	if x.SkuList == nil {
		return n
	}
	for i := range x.GetSkuList() {
		n += fastpb.SizeMessage(11, x.GetSkuList()[i])
	}
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField12() (n int) {
	if x.SpecList == nil {
		return n
	}
	for i := range x.GetSpecList() {
		n += fastpb.SizeMessage(12, x.GetSpecList()[i])
	}
	return n
}

func (x *GetGoodsListRsp_SpuInfo) sizeField13() (n int) {
	if x.SpecValueList == nil {
		return n
	}
	for i := range x.GetSpecValueList() {
		n += fastpb.SizeMessage(13, x.GetSpecValueList()[i])
	}
	return n
}

func (x *GetGoodsListRsp_SkuInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *GetGoodsListRsp_SkuInfo) sizeField1() (n int) {
	if x.SkuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSkuId())
	return n
}

func (x *GetGoodsListRsp_SkuInfo) sizeField2() (n int) {
	if x.SpuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetSpuId())
	return n
}

func (x *GetGoodsListRsp_SkuInfo) sizeField3() (n int) {
	if x.StockNum == 0 {
		return n
	}
	n += fastpb.SizeUint64(3, x.GetStockNum())
	return n
}

func (x *GetGoodsListRsp_SkuInfo) sizeField4() (n int) {
	if x.SkuAmount == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSkuAmount())
	return n
}

func (x *GetGoodsListRsp_SkuInfo) sizeField5() (n int) {
	if x.SkuName == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetSkuName())
	return n
}

func (x *GetGoodsListRsp_SkuInfo) sizeField6() (n int) {
	if x.SpecStr == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetSpecStr())
	return n
}

func (x *GetGoodsListRsp_SpecInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetGoodsListRsp_SpecInfo) sizeField1() (n int) {
	if x.SpecId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSpecId())
	return n
}

func (x *GetGoodsListRsp_SpecInfo) sizeField2() (n int) {
	if x.SpecName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSpecName())
	return n
}

func (x *GetGoodsListRsp_SpecInfo) sizeField3() (n int) {
	if x.SpuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(3, x.GetSpuId())
	return n
}

func (x *GetGoodsListRsp_SpecValueInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetGoodsListRsp_SpecValueInfo) sizeField1() (n int) {
	if x.SpecValueId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSpecValueId())
	return n
}

func (x *GetGoodsListRsp_SpecValueInfo) sizeField2() (n int) {
	if x.SpecName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSpecName())
	return n
}

func (x *GetGoodsListRsp_SpecValueInfo) sizeField3() (n int) {
	if x.SpecValue == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSpecValue())
	return n
}

var fieldIDToName_GetGoodsListReq = map[int32]string{
	1: "SpuIds",
}

var fieldIDToName_GetGoodsListRsp = map[int32]string{
	1: "Code",
	2: "Msg",
	3: "LogId",
	4: "Data",
}

var fieldIDToName_GetGoodsListRsp_GetGoodsListRspData = map[int32]string{
	1: "SpuList",
}

var fieldIDToName_GetGoodsListRsp_SpuInfo = map[int32]string{
	1:  "SpuId",
	2:  "SpuName",
	3:  "SpuDesc",
	4:  "SpuImgUrl",
	5:  "SpuMinAmount",
	6:  "SpuMaxAmount",
	7:  "SpuCategoryId",
	8:  "SpuCategoryName",
	9:  "BrandId",
	10: "BrandName",
	11: "SkuList",
	12: "SpecList",
	13: "SpecValueList",
}

var fieldIDToName_GetGoodsListRsp_SkuInfo = map[int32]string{
	1: "SkuId",
	2: "SpuId",
	3: "StockNum",
	4: "SkuAmount",
	5: "SkuName",
	6: "SpecStr",
}

var fieldIDToName_GetGoodsListRsp_SpecInfo = map[int32]string{
	1: "SpecId",
	2: "SpecName",
	3: "SpuId",
}

var fieldIDToName_GetGoodsListRsp_SpecValueInfo = map[int32]string{
	1: "SpecValueId",
	2: "SpecName",
	3: "SpecValue",
}
