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

func (x *BaseTrade) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseTrade[number], err)
}

func (x *BaseTrade) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TradeNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseTrade) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.TradeAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseTrade) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.BuyerId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseTrade) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SellerId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseTrade) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.PayType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrder) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseOrder[number], err)
}

func (x *BaseOrder) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrder) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.TradeNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrder) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ShopId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseOrder) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.OrderAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrder) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.SellerId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseOrder) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.BuyerId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseSpecValue) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseSpecValue[number], err)
}

func (x *BaseSpecValue) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseSpecValue) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpecId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseSpecValue) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SpecName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseSpecValue) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SpecValue, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 14:
		offset, err = x.fastReadField14(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 15:
		offset, err = x.fastReadField15(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 16:
		offset, err = x.fastReadField16(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseOrderItem[number], err)
}

func (x *BaseOrderItem) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SkuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SpuName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.CategoryId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.CategoryName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.BrandId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.BrandName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.SkuImgUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	x.SkuAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	var v BaseSpecValue
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SpecValueList = append(x.SpecValueList, &v)
	return offset, nil
}

func (x *BaseOrderItem) fastReadField12(buf []byte, _type int8) (offset int, err error) {
	x.TradeNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField13(buf []byte, _type int8) (offset int, err error) {
	x.OrderNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField14(buf []byte, _type int8) (offset int, err error) {
	x.ShopId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField15(buf []byte, _type int8) (offset int, err error) {
	x.SellerId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseOrderItem) fastReadField16(buf []byte, _type int8) (offset int, err error) {
	x.BuyerId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseTrade) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *BaseTrade) fastWriteField1(buf []byte) (offset int) {
	if x.TradeNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTradeNo())
	return offset
}

func (x *BaseTrade) fastWriteField2(buf []byte) (offset int) {
	if x.TradeAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTradeAmount())
	return offset
}

func (x *BaseTrade) fastWriteField3(buf []byte) (offset int) {
	if x.BuyerId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 3, x.GetBuyerId())
	return offset
}

func (x *BaseTrade) fastWriteField4(buf []byte) (offset int) {
	if x.SellerId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 4, x.GetSellerId())
	return offset
}

func (x *BaseTrade) fastWriteField5(buf []byte) (offset int) {
	if x.PayType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetPayType())
	return offset
}

func (x *BaseOrder) FastWrite(buf []byte) (offset int) {
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

func (x *BaseOrder) fastWriteField1(buf []byte) (offset int) {
	if x.OrderNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetOrderNo())
	return offset
}

func (x *BaseOrder) fastWriteField2(buf []byte) (offset int) {
	if x.TradeNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTradeNo())
	return offset
}

func (x *BaseOrder) fastWriteField3(buf []byte) (offset int) {
	if x.ShopId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 3, x.GetShopId())
	return offset
}

func (x *BaseOrder) fastWriteField4(buf []byte) (offset int) {
	if x.OrderAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetOrderAmount())
	return offset
}

func (x *BaseOrder) fastWriteField5(buf []byte) (offset int) {
	if x.SellerId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 5, x.GetSellerId())
	return offset
}

func (x *BaseOrder) fastWriteField6(buf []byte) (offset int) {
	if x.BuyerId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 6, x.GetBuyerId())
	return offset
}

func (x *BaseSpecValue) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *BaseSpecValue) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *BaseSpecValue) fastWriteField2(buf []byte) (offset int) {
	if x.SpecId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetSpecId())
	return offset
}

func (x *BaseSpecValue) fastWriteField3(buf []byte) (offset int) {
	if x.SpecName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSpecName())
	return offset
}

func (x *BaseSpecValue) fastWriteField4(buf []byte) (offset int) {
	if x.SpecValue == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSpecValue())
	return offset
}

func (x *BaseOrderItem) FastWrite(buf []byte) (offset int) {
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
	offset += x.fastWriteField14(buf[offset:])
	offset += x.fastWriteField15(buf[offset:])
	offset += x.fastWriteField16(buf[offset:])
	return offset
}

func (x *BaseOrderItem) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *BaseOrderItem) fastWriteField2(buf []byte) (offset int) {
	if x.SpuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetSpuId())
	return offset
}

func (x *BaseOrderItem) fastWriteField3(buf []byte) (offset int) {
	if x.SkuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 3, x.GetSkuId())
	return offset
}

func (x *BaseOrderItem) fastWriteField4(buf []byte) (offset int) {
	if x.SpuName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSpuName())
	return offset
}

func (x *BaseOrderItem) fastWriteField5(buf []byte) (offset int) {
	if x.CategoryId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 5, x.GetCategoryId())
	return offset
}

func (x *BaseOrderItem) fastWriteField6(buf []byte) (offset int) {
	if x.CategoryName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetCategoryName())
	return offset
}

func (x *BaseOrderItem) fastWriteField7(buf []byte) (offset int) {
	if x.BrandId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 7, x.GetBrandId())
	return offset
}

func (x *BaseOrderItem) fastWriteField8(buf []byte) (offset int) {
	if x.BrandName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetBrandName())
	return offset
}

func (x *BaseOrderItem) fastWriteField9(buf []byte) (offset int) {
	if x.SkuImgUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 9, x.GetSkuImgUrl())
	return offset
}

func (x *BaseOrderItem) fastWriteField10(buf []byte) (offset int) {
	if x.SkuAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 10, x.GetSkuAmount())
	return offset
}

func (x *BaseOrderItem) fastWriteField11(buf []byte) (offset int) {
	if x.SpecValueList == nil {
		return offset
	}
	for i := range x.GetSpecValueList() {
		offset += fastpb.WriteMessage(buf[offset:], 11, x.GetSpecValueList()[i])
	}
	return offset
}

func (x *BaseOrderItem) fastWriteField12(buf []byte) (offset int) {
	if x.TradeNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 12, x.GetTradeNo())
	return offset
}

func (x *BaseOrderItem) fastWriteField13(buf []byte) (offset int) {
	if x.OrderNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 13, x.GetOrderNo())
	return offset
}

func (x *BaseOrderItem) fastWriteField14(buf []byte) (offset int) {
	if x.ShopId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 14, x.GetShopId())
	return offset
}

func (x *BaseOrderItem) fastWriteField15(buf []byte) (offset int) {
	if x.SellerId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 15, x.GetSellerId())
	return offset
}

func (x *BaseOrderItem) fastWriteField16(buf []byte) (offset int) {
	if x.BuyerId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 16, x.GetBuyerId())
	return offset
}

func (x *BaseTrade) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *BaseTrade) sizeField1() (n int) {
	if x.TradeNo == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTradeNo())
	return n
}

func (x *BaseTrade) sizeField2() (n int) {
	if x.TradeAmount == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTradeAmount())
	return n
}

func (x *BaseTrade) sizeField3() (n int) {
	if x.BuyerId == 0 {
		return n
	}
	n += fastpb.SizeUint64(3, x.GetBuyerId())
	return n
}

func (x *BaseTrade) sizeField4() (n int) {
	if x.SellerId == 0 {
		return n
	}
	n += fastpb.SizeUint64(4, x.GetSellerId())
	return n
}

func (x *BaseTrade) sizeField5() (n int) {
	if x.PayType == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetPayType())
	return n
}

func (x *BaseOrder) Size() (n int) {
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

func (x *BaseOrder) sizeField1() (n int) {
	if x.OrderNo == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetOrderNo())
	return n
}

func (x *BaseOrder) sizeField2() (n int) {
	if x.TradeNo == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTradeNo())
	return n
}

func (x *BaseOrder) sizeField3() (n int) {
	if x.ShopId == 0 {
		return n
	}
	n += fastpb.SizeUint64(3, x.GetShopId())
	return n
}

func (x *BaseOrder) sizeField4() (n int) {
	if x.OrderAmount == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetOrderAmount())
	return n
}

func (x *BaseOrder) sizeField5() (n int) {
	if x.SellerId == 0 {
		return n
	}
	n += fastpb.SizeUint64(5, x.GetSellerId())
	return n
}

func (x *BaseOrder) sizeField6() (n int) {
	if x.BuyerId == 0 {
		return n
	}
	n += fastpb.SizeUint64(6, x.GetBuyerId())
	return n
}

func (x *BaseSpecValue) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *BaseSpecValue) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *BaseSpecValue) sizeField2() (n int) {
	if x.SpecId == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetSpecId())
	return n
}

func (x *BaseSpecValue) sizeField3() (n int) {
	if x.SpecName == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSpecName())
	return n
}

func (x *BaseSpecValue) sizeField4() (n int) {
	if x.SpecValue == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSpecValue())
	return n
}

func (x *BaseOrderItem) Size() (n int) {
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
	n += x.sizeField14()
	n += x.sizeField15()
	n += x.sizeField16()
	return n
}

func (x *BaseOrderItem) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *BaseOrderItem) sizeField2() (n int) {
	if x.SpuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetSpuId())
	return n
}

func (x *BaseOrderItem) sizeField3() (n int) {
	if x.SkuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(3, x.GetSkuId())
	return n
}

func (x *BaseOrderItem) sizeField4() (n int) {
	if x.SpuName == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSpuName())
	return n
}

func (x *BaseOrderItem) sizeField5() (n int) {
	if x.CategoryId == 0 {
		return n
	}
	n += fastpb.SizeUint64(5, x.GetCategoryId())
	return n
}

func (x *BaseOrderItem) sizeField6() (n int) {
	if x.CategoryName == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetCategoryName())
	return n
}

func (x *BaseOrderItem) sizeField7() (n int) {
	if x.BrandId == 0 {
		return n
	}
	n += fastpb.SizeUint64(7, x.GetBrandId())
	return n
}

func (x *BaseOrderItem) sizeField8() (n int) {
	if x.BrandName == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetBrandName())
	return n
}

func (x *BaseOrderItem) sizeField9() (n int) {
	if x.SkuImgUrl == "" {
		return n
	}
	n += fastpb.SizeString(9, x.GetSkuImgUrl())
	return n
}

func (x *BaseOrderItem) sizeField10() (n int) {
	if x.SkuAmount == "" {
		return n
	}
	n += fastpb.SizeString(10, x.GetSkuAmount())
	return n
}

func (x *BaseOrderItem) sizeField11() (n int) {
	if x.SpecValueList == nil {
		return n
	}
	for i := range x.GetSpecValueList() {
		n += fastpb.SizeMessage(11, x.GetSpecValueList()[i])
	}
	return n
}

func (x *BaseOrderItem) sizeField12() (n int) {
	if x.TradeNo == "" {
		return n
	}
	n += fastpb.SizeString(12, x.GetTradeNo())
	return n
}

func (x *BaseOrderItem) sizeField13() (n int) {
	if x.OrderNo == "" {
		return n
	}
	n += fastpb.SizeString(13, x.GetOrderNo())
	return n
}

func (x *BaseOrderItem) sizeField14() (n int) {
	if x.ShopId == 0 {
		return n
	}
	n += fastpb.SizeUint64(14, x.GetShopId())
	return n
}

func (x *BaseOrderItem) sizeField15() (n int) {
	if x.SellerId == 0 {
		return n
	}
	n += fastpb.SizeUint64(15, x.GetSellerId())
	return n
}

func (x *BaseOrderItem) sizeField16() (n int) {
	if x.BuyerId == 0 {
		return n
	}
	n += fastpb.SizeUint64(16, x.GetBuyerId())
	return n
}

var fieldIDToName_BaseTrade = map[int32]string{
	1: "TradeNo",
	2: "TradeAmount",
	3: "BuyerId",
	4: "SellerId",
	5: "PayType",
}

var fieldIDToName_BaseOrder = map[int32]string{
	1: "OrderNo",
	2: "TradeNo",
	3: "ShopId",
	4: "OrderAmount",
	5: "SellerId",
	6: "BuyerId",
}

var fieldIDToName_BaseSpecValue = map[int32]string{
	1: "Id",
	2: "SpecId",
	3: "SpecName",
	4: "SpecValue",
}

var fieldIDToName_BaseOrderItem = map[int32]string{
	1:  "Id",
	2:  "SpuId",
	3:  "SkuId",
	4:  "SpuName",
	5:  "CategoryId",
	6:  "CategoryName",
	7:  "BrandId",
	8:  "BrandName",
	9:  "SkuImgUrl",
	10: "SkuAmount",
	11: "SpecValueList",
	12: "TradeNo",
	13: "OrderNo",
	14: "ShopId",
	15: "SellerId",
	16: "BuyerId",
}
