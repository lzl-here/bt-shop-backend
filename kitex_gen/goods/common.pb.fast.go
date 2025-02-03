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

func (x *SpuInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SpuInfo[number], err)
}

func (x *SpuInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SpuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SpuInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpuName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpuInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SpuDesc, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpuInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SpuImgUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpuInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.SpuPrice, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpuInfo) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.SpuCategoryId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SpuInfo) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.SpuCategoryName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpuInfo) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.BrandId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SpuInfo) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	x.BrandName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpuInfo) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	var v SkuInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SkuList = append(x.SkuList, &v)
	return offset, nil
}

func (x *SpuInfo) fastReadField12(buf []byte, _type int8) (offset int, err error) {
	var v SpecInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SpecList = append(x.SpecList, &v)
	return offset, nil
}

func (x *SpuInfo) fastReadField13(buf []byte, _type int8) (offset int, err error) {
	var v SpecValueInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SpecValueList = append(x.SpecValueList, &v)
	return offset, nil
}

func (x *SkuInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SkuInfo[number], err)
}

func (x *SkuInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SkuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SkuInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SkuInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.StockNum, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SkuInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SkuAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SkuInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.SkuName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SkuInfo) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.SpecStr, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpecInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SpecInfo[number], err)
}

func (x *SpecInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SpecId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SpecInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpecName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpecInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SpuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SpecValueInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SpecValueInfo[number], err)
}

func (x *SpecValueInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SpecValueId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SpecValueInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpecName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpecValueInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SpecValue, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpecKeyValue) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SpecKeyValue[number], err)
}

func (x *SpecKeyValue) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SpecName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpecKeyValue) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SpecValue, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AttributeInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AttributeInfo[number], err)
}

func (x *AttributeInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *AttributeInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.AttributeName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AttributeInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.AttributeValue, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Keyword) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Keyword[number], err)
}

func (x *Keyword) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *Keyword) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Value, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Category) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Category[number], err)
}

func (x *Category) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.CategoryId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *Category) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CategoryName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Category) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ParentId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *Category) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v Category
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Children = append(x.Children, &v)
	return offset, nil
}

func (x *Brand) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Brand[number], err)
}

func (x *Brand) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.BrandId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *Brand) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.BrandName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SpuInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	offset += x.fastWriteField12(buf[offset:])
	offset += x.fastWriteField13(buf[offset:])
	return offset
}

func (x *SpuInfo) fastWriteField1(buf []byte) (offset int) {
	if x.SpuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSpuId())
	return offset
}

func (x *SpuInfo) fastWriteField2(buf []byte) (offset int) {
	if x.SpuName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSpuName())
	return offset
}

func (x *SpuInfo) fastWriteField3(buf []byte) (offset int) {
	if x.SpuDesc == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSpuDesc())
	return offset
}

func (x *SpuInfo) fastWriteField4(buf []byte) (offset int) {
	if x.SpuImgUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSpuImgUrl())
	return offset
}

func (x *SpuInfo) fastWriteField5(buf []byte) (offset int) {
	if x.SpuPrice == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetSpuPrice())
	return offset
}

func (x *SpuInfo) fastWriteField7(buf []byte) (offset int) {
	if x.SpuCategoryId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 7, x.GetSpuCategoryId())
	return offset
}

func (x *SpuInfo) fastWriteField8(buf []byte) (offset int) {
	if x.SpuCategoryName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetSpuCategoryName())
	return offset
}

func (x *SpuInfo) fastWriteField9(buf []byte) (offset int) {
	if x.BrandId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 9, x.GetBrandId())
	return offset
}

func (x *SpuInfo) fastWriteField10(buf []byte) (offset int) {
	if x.BrandName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 10, x.GetBrandName())
	return offset
}

func (x *SpuInfo) fastWriteField11(buf []byte) (offset int) {
	if x.SkuList == nil {
		return offset
	}
	for i := range x.GetSkuList() {
		offset += fastpb.WriteMessage(buf[offset:], 11, x.GetSkuList()[i])
	}
	return offset
}

func (x *SpuInfo) fastWriteField12(buf []byte) (offset int) {
	if x.SpecList == nil {
		return offset
	}
	for i := range x.GetSpecList() {
		offset += fastpb.WriteMessage(buf[offset:], 12, x.GetSpecList()[i])
	}
	return offset
}

func (x *SpuInfo) fastWriteField13(buf []byte) (offset int) {
	if x.SpecValueList == nil {
		return offset
	}
	for i := range x.GetSpecValueList() {
		offset += fastpb.WriteMessage(buf[offset:], 13, x.GetSpecValueList()[i])
	}
	return offset
}

func (x *SkuInfo) FastWrite(buf []byte) (offset int) {
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

func (x *SkuInfo) fastWriteField1(buf []byte) (offset int) {
	if x.SkuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSkuId())
	return offset
}

func (x *SkuInfo) fastWriteField2(buf []byte) (offset int) {
	if x.SpuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetSpuId())
	return offset
}

func (x *SkuInfo) fastWriteField3(buf []byte) (offset int) {
	if x.StockNum == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 3, x.GetStockNum())
	return offset
}

func (x *SkuInfo) fastWriteField4(buf []byte) (offset int) {
	if x.SkuAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSkuAmount())
	return offset
}

func (x *SkuInfo) fastWriteField5(buf []byte) (offset int) {
	if x.SkuName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetSkuName())
	return offset
}

func (x *SkuInfo) fastWriteField6(buf []byte) (offset int) {
	if x.SpecStr == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetSpecStr())
	return offset
}

func (x *SpecInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *SpecInfo) fastWriteField1(buf []byte) (offset int) {
	if x.SpecId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSpecId())
	return offset
}

func (x *SpecInfo) fastWriteField2(buf []byte) (offset int) {
	if x.SpecName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSpecName())
	return offset
}

func (x *SpecInfo) fastWriteField3(buf []byte) (offset int) {
	if x.SpuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 3, x.GetSpuId())
	return offset
}

func (x *SpecValueInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *SpecValueInfo) fastWriteField1(buf []byte) (offset int) {
	if x.SpecValueId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSpecValueId())
	return offset
}

func (x *SpecValueInfo) fastWriteField2(buf []byte) (offset int) {
	if x.SpecName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSpecName())
	return offset
}

func (x *SpecValueInfo) fastWriteField3(buf []byte) (offset int) {
	if x.SpecValue == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSpecValue())
	return offset
}

func (x *SpecKeyValue) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SpecKeyValue) fastWriteField1(buf []byte) (offset int) {
	if x.SpecName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSpecName())
	return offset
}

func (x *SpecKeyValue) fastWriteField2(buf []byte) (offset int) {
	if x.SpecValue == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSpecValue())
	return offset
}

func (x *AttributeInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *AttributeInfo) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *AttributeInfo) fastWriteField2(buf []byte) (offset int) {
	if x.AttributeName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetAttributeName())
	return offset
}

func (x *AttributeInfo) fastWriteField3(buf []byte) (offset int) {
	if x.AttributeValue == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAttributeValue())
	return offset
}

func (x *Keyword) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *Keyword) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Keyword) fastWriteField2(buf []byte) (offset int) {
	if x.Value == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetValue())
	return offset
}

func (x *Category) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *Category) fastWriteField1(buf []byte) (offset int) {
	if x.CategoryId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetCategoryId())
	return offset
}

func (x *Category) fastWriteField2(buf []byte) (offset int) {
	if x.CategoryName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCategoryName())
	return offset
}

func (x *Category) fastWriteField3(buf []byte) (offset int) {
	if x.ParentId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 3, x.GetParentId())
	return offset
}

func (x *Category) fastWriteField4(buf []byte) (offset int) {
	if x.Children == nil {
		return offset
	}
	for i := range x.GetChildren() {
		offset += fastpb.WriteMessage(buf[offset:], 4, x.GetChildren()[i])
	}
	return offset
}

func (x *Brand) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *Brand) fastWriteField1(buf []byte) (offset int) {
	if x.BrandId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetBrandId())
	return offset
}

func (x *Brand) fastWriteField2(buf []byte) (offset int) {
	if x.BrandName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetBrandName())
	return offset
}

func (x *SpuInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	n += x.sizeField12()
	n += x.sizeField13()
	return n
}

func (x *SpuInfo) sizeField1() (n int) {
	if x.SpuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSpuId())
	return n
}

func (x *SpuInfo) sizeField2() (n int) {
	if x.SpuName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSpuName())
	return n
}

func (x *SpuInfo) sizeField3() (n int) {
	if x.SpuDesc == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSpuDesc())
	return n
}

func (x *SpuInfo) sizeField4() (n int) {
	if x.SpuImgUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSpuImgUrl())
	return n
}

func (x *SpuInfo) sizeField5() (n int) {
	if x.SpuPrice == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetSpuPrice())
	return n
}

func (x *SpuInfo) sizeField7() (n int) {
	if x.SpuCategoryId == 0 {
		return n
	}
	n += fastpb.SizeUint64(7, x.GetSpuCategoryId())
	return n
}

func (x *SpuInfo) sizeField8() (n int) {
	if x.SpuCategoryName == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetSpuCategoryName())
	return n
}

func (x *SpuInfo) sizeField9() (n int) {
	if x.BrandId == 0 {
		return n
	}
	n += fastpb.SizeUint64(9, x.GetBrandId())
	return n
}

func (x *SpuInfo) sizeField10() (n int) {
	if x.BrandName == "" {
		return n
	}
	n += fastpb.SizeString(10, x.GetBrandName())
	return n
}

func (x *SpuInfo) sizeField11() (n int) {
	if x.SkuList == nil {
		return n
	}
	for i := range x.GetSkuList() {
		n += fastpb.SizeMessage(11, x.GetSkuList()[i])
	}
	return n
}

func (x *SpuInfo) sizeField12() (n int) {
	if x.SpecList == nil {
		return n
	}
	for i := range x.GetSpecList() {
		n += fastpb.SizeMessage(12, x.GetSpecList()[i])
	}
	return n
}

func (x *SpuInfo) sizeField13() (n int) {
	if x.SpecValueList == nil {
		return n
	}
	for i := range x.GetSpecValueList() {
		n += fastpb.SizeMessage(13, x.GetSpecValueList()[i])
	}
	return n
}

func (x *SkuInfo) Size() (n int) {
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

func (x *SkuInfo) sizeField1() (n int) {
	if x.SkuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSkuId())
	return n
}

func (x *SkuInfo) sizeField2() (n int) {
	if x.SpuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetSpuId())
	return n
}

func (x *SkuInfo) sizeField3() (n int) {
	if x.StockNum == 0 {
		return n
	}
	n += fastpb.SizeUint64(3, x.GetStockNum())
	return n
}

func (x *SkuInfo) sizeField4() (n int) {
	if x.SkuAmount == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSkuAmount())
	return n
}

func (x *SkuInfo) sizeField5() (n int) {
	if x.SkuName == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetSkuName())
	return n
}

func (x *SkuInfo) sizeField6() (n int) {
	if x.SpecStr == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetSpecStr())
	return n
}

func (x *SpecInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *SpecInfo) sizeField1() (n int) {
	if x.SpecId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSpecId())
	return n
}

func (x *SpecInfo) sizeField2() (n int) {
	if x.SpecName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSpecName())
	return n
}

func (x *SpecInfo) sizeField3() (n int) {
	if x.SpuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(3, x.GetSpuId())
	return n
}

func (x *SpecValueInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *SpecValueInfo) sizeField1() (n int) {
	if x.SpecValueId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSpecValueId())
	return n
}

func (x *SpecValueInfo) sizeField2() (n int) {
	if x.SpecName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSpecName())
	return n
}

func (x *SpecValueInfo) sizeField3() (n int) {
	if x.SpecValue == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSpecValue())
	return n
}

func (x *SpecKeyValue) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SpecKeyValue) sizeField1() (n int) {
	if x.SpecName == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSpecName())
	return n
}

func (x *SpecKeyValue) sizeField2() (n int) {
	if x.SpecValue == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSpecValue())
	return n
}

func (x *AttributeInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *AttributeInfo) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *AttributeInfo) sizeField2() (n int) {
	if x.AttributeName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetAttributeName())
	return n
}

func (x *AttributeInfo) sizeField3() (n int) {
	if x.AttributeValue == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAttributeValue())
	return n
}

func (x *Keyword) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *Keyword) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *Keyword) sizeField2() (n int) {
	if x.Value == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetValue())
	return n
}

func (x *Category) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *Category) sizeField1() (n int) {
	if x.CategoryId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetCategoryId())
	return n
}

func (x *Category) sizeField2() (n int) {
	if x.CategoryName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetCategoryName())
	return n
}

func (x *Category) sizeField3() (n int) {
	if x.ParentId == 0 {
		return n
	}
	n += fastpb.SizeUint64(3, x.GetParentId())
	return n
}

func (x *Category) sizeField4() (n int) {
	if x.Children == nil {
		return n
	}
	for i := range x.GetChildren() {
		n += fastpb.SizeMessage(4, x.GetChildren()[i])
	}
	return n
}

func (x *Brand) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *Brand) sizeField1() (n int) {
	if x.BrandId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetBrandId())
	return n
}

func (x *Brand) sizeField2() (n int) {
	if x.BrandName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetBrandName())
	return n
}

var fieldIDToName_SpuInfo = map[int32]string{
	1:  "SpuId",
	2:  "SpuName",
	3:  "SpuDesc",
	4:  "SpuImgUrl",
	5:  "SpuPrice",
	7:  "SpuCategoryId",
	8:  "SpuCategoryName",
	9:  "BrandId",
	10: "BrandName",
	11: "SkuList",
	12: "SpecList",
	13: "SpecValueList",
}

var fieldIDToName_SkuInfo = map[int32]string{
	1: "SkuId",
	2: "SpuId",
	3: "StockNum",
	4: "SkuAmount",
	5: "SkuName",
	6: "SpecStr",
}

var fieldIDToName_SpecInfo = map[int32]string{
	1: "SpecId",
	2: "SpecName",
	3: "SpuId",
}

var fieldIDToName_SpecValueInfo = map[int32]string{
	1: "SpecValueId",
	2: "SpecName",
	3: "SpecValue",
}

var fieldIDToName_SpecKeyValue = map[int32]string{
	1: "SpecName",
	2: "SpecValue",
}

var fieldIDToName_AttributeInfo = map[int32]string{
	1: "Id",
	2: "AttributeName",
	3: "AttributeValue",
}

var fieldIDToName_Keyword = map[int32]string{
	1: "Id",
	2: "Value",
}

var fieldIDToName_Category = map[int32]string{
	1: "CategoryId",
	2: "CategoryName",
	3: "ParentId",
	4: "Children",
}

var fieldIDToName_Brand = map[int32]string{
	1: "BrandId",
	2: "BrandName",
}
