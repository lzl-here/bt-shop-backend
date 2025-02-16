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

func (x *AlipayWebhookReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 17:
		offset, err = x.fastReadField17(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 18:
		offset, err = x.fastReadField18(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 19:
		offset, err = x.fastReadField19(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 20:
		offset, err = x.fastReadField20(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 21:
		offset, err = x.fastReadField21(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 22:
		offset, err = x.fastReadField22(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 23:
		offset, err = x.fastReadField23(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 24:
		offset, err = x.fastReadField24(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 25:
		offset, err = x.fastReadField25(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 26:
		offset, err = x.fastReadField26(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 27:
		offset, err = x.fastReadField27(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 28:
		offset, err = x.fastReadField28(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 29:
		offset, err = x.fastReadField29(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 30:
		offset, err = x.fastReadField30(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 31:
		offset, err = x.fastReadField31(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 32:
		offset, err = x.fastReadField32(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 33:
		offset, err = x.fastReadField33(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 34:
		offset, err = x.fastReadField34(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 35:
		offset, err = x.fastReadField35(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 36:
		offset, err = x.fastReadField36(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 37:
		offset, err = x.fastReadField37(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 38:
		offset, err = x.fastReadField38(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 39:
		offset, err = x.fastReadField39(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 40:
		offset, err = x.fastReadField40(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 41:
		offset, err = x.fastReadField41(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 42:
		offset, err = x.fastReadField42(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlipayWebhookReq[number], err)
}

func (x *AlipayWebhookReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.AuthAppId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.NotifyTime, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.NotifyType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.NotifyId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.AppId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Charset, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.SignType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.Sign, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	x.TradeNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	x.OutTradeNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField12(buf []byte, _type int8) (offset int, err error) {
	x.OutRequestNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField13(buf []byte, _type int8) (offset int, err error) {
	x.OutBizNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField14(buf []byte, _type int8) (offset int, err error) {
	x.BuyerId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField15(buf []byte, _type int8) (offset int, err error) {
	x.BuyerLogonId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField16(buf []byte, _type int8) (offset int, err error) {
	x.BuyerOpenId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField17(buf []byte, _type int8) (offset int, err error) {
	x.SellerId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField18(buf []byte, _type int8) (offset int, err error) {
	x.SellerEmail, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField19(buf []byte, _type int8) (offset int, err error) {
	x.TradeStatus, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField20(buf []byte, _type int8) (offset int, err error) {
	x.RefundStatus, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField21(buf []byte, _type int8) (offset int, err error) {
	x.RefundReason, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField22(buf []byte, _type int8) (offset int, err error) {
	x.RefundAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField23(buf []byte, _type int8) (offset int, err error) {
	x.TotalAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField24(buf []byte, _type int8) (offset int, err error) {
	x.ReceiptAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField25(buf []byte, _type int8) (offset int, err error) {
	x.InvoiceAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField26(buf []byte, _type int8) (offset int, err error) {
	x.BuyerPayAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField27(buf []byte, _type int8) (offset int, err error) {
	x.PointAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField28(buf []byte, _type int8) (offset int, err error) {
	x.RefundFee, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField29(buf []byte, _type int8) (offset int, err error) {
	x.Subject, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField30(buf []byte, _type int8) (offset int, err error) {
	x.Body, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField31(buf []byte, _type int8) (offset int, err error) {
	x.GmtCreate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField32(buf []byte, _type int8) (offset int, err error) {
	x.GmtPayment, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField33(buf []byte, _type int8) (offset int, err error) {
	x.GmtRefund, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField34(buf []byte, _type int8) (offset int, err error) {
	x.GmtClose, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField35(buf []byte, _type int8) (offset int, err error) {
	x.FundBillList, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField36(buf []byte, _type int8) (offset int, err error) {
	x.PassbackParams, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField37(buf []byte, _type int8) (offset int, err error) {
	x.VoucherDetailList, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField38(buf []byte, _type int8) (offset int, err error) {
	x.AgreementNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField39(buf []byte, _type int8) (offset int, err error) {
	x.ExternalAgreementNo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField40(buf []byte, _type int8) (offset int, err error) {
	x.DbackStatus, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField41(buf []byte, _type int8) (offset int, err error) {
	x.DbackAmount, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookReq) fastReadField42(buf []byte, _type int8) (offset int, err error) {
	x.BankAckTime, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookRsp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlipayWebhookRsp[number], err)
}

func (x *AlipayWebhookRsp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AlipayWebhookRsp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookRsp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.LogId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlipayWebhookRsp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v AlipayWebhookRsp_AlipayWebhookRspData
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Data = &v
	return offset, nil
}

func (x *AlipayWebhookRsp_AlipayWebhookRspData) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *AlipayWebhookReq) FastWrite(buf []byte) (offset int) {
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
	offset += x.fastWriteField17(buf[offset:])
	offset += x.fastWriteField18(buf[offset:])
	offset += x.fastWriteField19(buf[offset:])
	offset += x.fastWriteField20(buf[offset:])
	offset += x.fastWriteField21(buf[offset:])
	offset += x.fastWriteField22(buf[offset:])
	offset += x.fastWriteField23(buf[offset:])
	offset += x.fastWriteField24(buf[offset:])
	offset += x.fastWriteField25(buf[offset:])
	offset += x.fastWriteField26(buf[offset:])
	offset += x.fastWriteField27(buf[offset:])
	offset += x.fastWriteField28(buf[offset:])
	offset += x.fastWriteField29(buf[offset:])
	offset += x.fastWriteField30(buf[offset:])
	offset += x.fastWriteField31(buf[offset:])
	offset += x.fastWriteField32(buf[offset:])
	offset += x.fastWriteField33(buf[offset:])
	offset += x.fastWriteField34(buf[offset:])
	offset += x.fastWriteField35(buf[offset:])
	offset += x.fastWriteField36(buf[offset:])
	offset += x.fastWriteField37(buf[offset:])
	offset += x.fastWriteField38(buf[offset:])
	offset += x.fastWriteField39(buf[offset:])
	offset += x.fastWriteField40(buf[offset:])
	offset += x.fastWriteField41(buf[offset:])
	offset += x.fastWriteField42(buf[offset:])
	return offset
}

func (x *AlipayWebhookReq) fastWriteField1(buf []byte) (offset int) {
	if x.AuthAppId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAuthAppId())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField2(buf []byte) (offset int) {
	if x.NotifyTime == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetNotifyTime())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField3(buf []byte) (offset int) {
	if x.NotifyType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetNotifyType())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField4(buf []byte) (offset int) {
	if x.NotifyId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetNotifyId())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField5(buf []byte) (offset int) {
	if x.AppId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetAppId())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField6(buf []byte) (offset int) {
	if x.Charset == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetCharset())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField7(buf []byte) (offset int) {
	if x.Version == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetVersion())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField8(buf []byte) (offset int) {
	if x.SignType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetSignType())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField9(buf []byte) (offset int) {
	if x.Sign == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 9, x.GetSign())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField10(buf []byte) (offset int) {
	if x.TradeNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 10, x.GetTradeNo())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField11(buf []byte) (offset int) {
	if x.OutTradeNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 11, x.GetOutTradeNo())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField12(buf []byte) (offset int) {
	if x.OutRequestNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 12, x.GetOutRequestNo())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField13(buf []byte) (offset int) {
	if x.OutBizNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 13, x.GetOutBizNo())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField14(buf []byte) (offset int) {
	if x.BuyerId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 14, x.GetBuyerId())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField15(buf []byte) (offset int) {
	if x.BuyerLogonId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 15, x.GetBuyerLogonId())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField16(buf []byte) (offset int) {
	if x.BuyerOpenId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 16, x.GetBuyerOpenId())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField17(buf []byte) (offset int) {
	if x.SellerId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 17, x.GetSellerId())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField18(buf []byte) (offset int) {
	if x.SellerEmail == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 18, x.GetSellerEmail())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField19(buf []byte) (offset int) {
	if x.TradeStatus == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 19, x.GetTradeStatus())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField20(buf []byte) (offset int) {
	if x.RefundStatus == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 20, x.GetRefundStatus())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField21(buf []byte) (offset int) {
	if x.RefundReason == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 21, x.GetRefundReason())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField22(buf []byte) (offset int) {
	if x.RefundAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 22, x.GetRefundAmount())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField23(buf []byte) (offset int) {
	if x.TotalAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 23, x.GetTotalAmount())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField24(buf []byte) (offset int) {
	if x.ReceiptAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 24, x.GetReceiptAmount())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField25(buf []byte) (offset int) {
	if x.InvoiceAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 25, x.GetInvoiceAmount())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField26(buf []byte) (offset int) {
	if x.BuyerPayAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 26, x.GetBuyerPayAmount())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField27(buf []byte) (offset int) {
	if x.PointAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 27, x.GetPointAmount())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField28(buf []byte) (offset int) {
	if x.RefundFee == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 28, x.GetRefundFee())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField29(buf []byte) (offset int) {
	if x.Subject == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 29, x.GetSubject())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField30(buf []byte) (offset int) {
	if x.Body == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 30, x.GetBody())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField31(buf []byte) (offset int) {
	if x.GmtCreate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 31, x.GetGmtCreate())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField32(buf []byte) (offset int) {
	if x.GmtPayment == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 32, x.GetGmtPayment())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField33(buf []byte) (offset int) {
	if x.GmtRefund == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 33, x.GetGmtRefund())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField34(buf []byte) (offset int) {
	if x.GmtClose == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 34, x.GetGmtClose())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField35(buf []byte) (offset int) {
	if x.FundBillList == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 35, x.GetFundBillList())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField36(buf []byte) (offset int) {
	if x.PassbackParams == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 36, x.GetPassbackParams())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField37(buf []byte) (offset int) {
	if x.VoucherDetailList == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 37, x.GetVoucherDetailList())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField38(buf []byte) (offset int) {
	if x.AgreementNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 38, x.GetAgreementNo())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField39(buf []byte) (offset int) {
	if x.ExternalAgreementNo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 39, x.GetExternalAgreementNo())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField40(buf []byte) (offset int) {
	if x.DbackStatus == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 40, x.GetDbackStatus())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField41(buf []byte) (offset int) {
	if x.DbackAmount == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 41, x.GetDbackAmount())
	return offset
}

func (x *AlipayWebhookReq) fastWriteField42(buf []byte) (offset int) {
	if x.BankAckTime == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 42, x.GetBankAckTime())
	return offset
}

func (x *AlipayWebhookRsp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *AlipayWebhookRsp) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *AlipayWebhookRsp) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *AlipayWebhookRsp) fastWriteField3(buf []byte) (offset int) {
	if x.LogId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetLogId())
	return offset
}

func (x *AlipayWebhookRsp) fastWriteField4(buf []byte) (offset int) {
	if x.Data == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetData())
	return offset
}

func (x *AlipayWebhookRsp_AlipayWebhookRspData) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *AlipayWebhookReq) Size() (n int) {
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
	n += x.sizeField17()
	n += x.sizeField18()
	n += x.sizeField19()
	n += x.sizeField20()
	n += x.sizeField21()
	n += x.sizeField22()
	n += x.sizeField23()
	n += x.sizeField24()
	n += x.sizeField25()
	n += x.sizeField26()
	n += x.sizeField27()
	n += x.sizeField28()
	n += x.sizeField29()
	n += x.sizeField30()
	n += x.sizeField31()
	n += x.sizeField32()
	n += x.sizeField33()
	n += x.sizeField34()
	n += x.sizeField35()
	n += x.sizeField36()
	n += x.sizeField37()
	n += x.sizeField38()
	n += x.sizeField39()
	n += x.sizeField40()
	n += x.sizeField41()
	n += x.sizeField42()
	return n
}

func (x *AlipayWebhookReq) sizeField1() (n int) {
	if x.AuthAppId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAuthAppId())
	return n
}

func (x *AlipayWebhookReq) sizeField2() (n int) {
	if x.NotifyTime == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetNotifyTime())
	return n
}

func (x *AlipayWebhookReq) sizeField3() (n int) {
	if x.NotifyType == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetNotifyType())
	return n
}

func (x *AlipayWebhookReq) sizeField4() (n int) {
	if x.NotifyId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetNotifyId())
	return n
}

func (x *AlipayWebhookReq) sizeField5() (n int) {
	if x.AppId == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetAppId())
	return n
}

func (x *AlipayWebhookReq) sizeField6() (n int) {
	if x.Charset == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetCharset())
	return n
}

func (x *AlipayWebhookReq) sizeField7() (n int) {
	if x.Version == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetVersion())
	return n
}

func (x *AlipayWebhookReq) sizeField8() (n int) {
	if x.SignType == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetSignType())
	return n
}

func (x *AlipayWebhookReq) sizeField9() (n int) {
	if x.Sign == "" {
		return n
	}
	n += fastpb.SizeString(9, x.GetSign())
	return n
}

func (x *AlipayWebhookReq) sizeField10() (n int) {
	if x.TradeNo == "" {
		return n
	}
	n += fastpb.SizeString(10, x.GetTradeNo())
	return n
}

func (x *AlipayWebhookReq) sizeField11() (n int) {
	if x.OutTradeNo == "" {
		return n
	}
	n += fastpb.SizeString(11, x.GetOutTradeNo())
	return n
}

func (x *AlipayWebhookReq) sizeField12() (n int) {
	if x.OutRequestNo == "" {
		return n
	}
	n += fastpb.SizeString(12, x.GetOutRequestNo())
	return n
}

func (x *AlipayWebhookReq) sizeField13() (n int) {
	if x.OutBizNo == "" {
		return n
	}
	n += fastpb.SizeString(13, x.GetOutBizNo())
	return n
}

func (x *AlipayWebhookReq) sizeField14() (n int) {
	if x.BuyerId == "" {
		return n
	}
	n += fastpb.SizeString(14, x.GetBuyerId())
	return n
}

func (x *AlipayWebhookReq) sizeField15() (n int) {
	if x.BuyerLogonId == "" {
		return n
	}
	n += fastpb.SizeString(15, x.GetBuyerLogonId())
	return n
}

func (x *AlipayWebhookReq) sizeField16() (n int) {
	if x.BuyerOpenId == "" {
		return n
	}
	n += fastpb.SizeString(16, x.GetBuyerOpenId())
	return n
}

func (x *AlipayWebhookReq) sizeField17() (n int) {
	if x.SellerId == "" {
		return n
	}
	n += fastpb.SizeString(17, x.GetSellerId())
	return n
}

func (x *AlipayWebhookReq) sizeField18() (n int) {
	if x.SellerEmail == "" {
		return n
	}
	n += fastpb.SizeString(18, x.GetSellerEmail())
	return n
}

func (x *AlipayWebhookReq) sizeField19() (n int) {
	if x.TradeStatus == "" {
		return n
	}
	n += fastpb.SizeString(19, x.GetTradeStatus())
	return n
}

func (x *AlipayWebhookReq) sizeField20() (n int) {
	if x.RefundStatus == "" {
		return n
	}
	n += fastpb.SizeString(20, x.GetRefundStatus())
	return n
}

func (x *AlipayWebhookReq) sizeField21() (n int) {
	if x.RefundReason == "" {
		return n
	}
	n += fastpb.SizeString(21, x.GetRefundReason())
	return n
}

func (x *AlipayWebhookReq) sizeField22() (n int) {
	if x.RefundAmount == "" {
		return n
	}
	n += fastpb.SizeString(22, x.GetRefundAmount())
	return n
}

func (x *AlipayWebhookReq) sizeField23() (n int) {
	if x.TotalAmount == "" {
		return n
	}
	n += fastpb.SizeString(23, x.GetTotalAmount())
	return n
}

func (x *AlipayWebhookReq) sizeField24() (n int) {
	if x.ReceiptAmount == "" {
		return n
	}
	n += fastpb.SizeString(24, x.GetReceiptAmount())
	return n
}

func (x *AlipayWebhookReq) sizeField25() (n int) {
	if x.InvoiceAmount == "" {
		return n
	}
	n += fastpb.SizeString(25, x.GetInvoiceAmount())
	return n
}

func (x *AlipayWebhookReq) sizeField26() (n int) {
	if x.BuyerPayAmount == "" {
		return n
	}
	n += fastpb.SizeString(26, x.GetBuyerPayAmount())
	return n
}

func (x *AlipayWebhookReq) sizeField27() (n int) {
	if x.PointAmount == "" {
		return n
	}
	n += fastpb.SizeString(27, x.GetPointAmount())
	return n
}

func (x *AlipayWebhookReq) sizeField28() (n int) {
	if x.RefundFee == "" {
		return n
	}
	n += fastpb.SizeString(28, x.GetRefundFee())
	return n
}

func (x *AlipayWebhookReq) sizeField29() (n int) {
	if x.Subject == "" {
		return n
	}
	n += fastpb.SizeString(29, x.GetSubject())
	return n
}

func (x *AlipayWebhookReq) sizeField30() (n int) {
	if x.Body == "" {
		return n
	}
	n += fastpb.SizeString(30, x.GetBody())
	return n
}

func (x *AlipayWebhookReq) sizeField31() (n int) {
	if x.GmtCreate == "" {
		return n
	}
	n += fastpb.SizeString(31, x.GetGmtCreate())
	return n
}

func (x *AlipayWebhookReq) sizeField32() (n int) {
	if x.GmtPayment == "" {
		return n
	}
	n += fastpb.SizeString(32, x.GetGmtPayment())
	return n
}

func (x *AlipayWebhookReq) sizeField33() (n int) {
	if x.GmtRefund == "" {
		return n
	}
	n += fastpb.SizeString(33, x.GetGmtRefund())
	return n
}

func (x *AlipayWebhookReq) sizeField34() (n int) {
	if x.GmtClose == "" {
		return n
	}
	n += fastpb.SizeString(34, x.GetGmtClose())
	return n
}

func (x *AlipayWebhookReq) sizeField35() (n int) {
	if x.FundBillList == "" {
		return n
	}
	n += fastpb.SizeString(35, x.GetFundBillList())
	return n
}

func (x *AlipayWebhookReq) sizeField36() (n int) {
	if x.PassbackParams == "" {
		return n
	}
	n += fastpb.SizeString(36, x.GetPassbackParams())
	return n
}

func (x *AlipayWebhookReq) sizeField37() (n int) {
	if x.VoucherDetailList == "" {
		return n
	}
	n += fastpb.SizeString(37, x.GetVoucherDetailList())
	return n
}

func (x *AlipayWebhookReq) sizeField38() (n int) {
	if x.AgreementNo == "" {
		return n
	}
	n += fastpb.SizeString(38, x.GetAgreementNo())
	return n
}

func (x *AlipayWebhookReq) sizeField39() (n int) {
	if x.ExternalAgreementNo == "" {
		return n
	}
	n += fastpb.SizeString(39, x.GetExternalAgreementNo())
	return n
}

func (x *AlipayWebhookReq) sizeField40() (n int) {
	if x.DbackStatus == "" {
		return n
	}
	n += fastpb.SizeString(40, x.GetDbackStatus())
	return n
}

func (x *AlipayWebhookReq) sizeField41() (n int) {
	if x.DbackAmount == "" {
		return n
	}
	n += fastpb.SizeString(41, x.GetDbackAmount())
	return n
}

func (x *AlipayWebhookReq) sizeField42() (n int) {
	if x.BankAckTime == "" {
		return n
	}
	n += fastpb.SizeString(42, x.GetBankAckTime())
	return n
}

func (x *AlipayWebhookRsp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *AlipayWebhookRsp) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetCode())
	return n
}

func (x *AlipayWebhookRsp) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *AlipayWebhookRsp) sizeField3() (n int) {
	if x.LogId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetLogId())
	return n
}

func (x *AlipayWebhookRsp) sizeField4() (n int) {
	if x.Data == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetData())
	return n
}

func (x *AlipayWebhookRsp_AlipayWebhookRspData) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_AlipayWebhookReq = map[int32]string{
	1:  "AuthAppId",
	2:  "NotifyTime",
	3:  "NotifyType",
	4:  "NotifyId",
	5:  "AppId",
	6:  "Charset",
	7:  "Version",
	8:  "SignType",
	9:  "Sign",
	10: "TradeNo",
	11: "OutTradeNo",
	12: "OutRequestNo",
	13: "OutBizNo",
	14: "BuyerId",
	15: "BuyerLogonId",
	16: "BuyerOpenId",
	17: "SellerId",
	18: "SellerEmail",
	19: "TradeStatus",
	20: "RefundStatus",
	21: "RefundReason",
	22: "RefundAmount",
	23: "TotalAmount",
	24: "ReceiptAmount",
	25: "InvoiceAmount",
	26: "BuyerPayAmount",
	27: "PointAmount",
	28: "RefundFee",
	29: "Subject",
	30: "Body",
	31: "GmtCreate",
	32: "GmtPayment",
	33: "GmtRefund",
	34: "GmtClose",
	35: "FundBillList",
	36: "PassbackParams",
	37: "VoucherDetailList",
	38: "AgreementNo",
	39: "ExternalAgreementNo",
	40: "DbackStatus",
	41: "DbackAmount",
	42: "BankAckTime",
}

var fieldIDToName_AlipayWebhookRsp = map[int32]string{
	1: "Code",
	2: "Msg",
	3: "LogId",
	4: "Data",
}

var fieldIDToName_AlipayWebhookRsp_AlipayWebhookRspData = map[int32]string{}
