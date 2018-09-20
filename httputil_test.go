package shouqianba

import (
	"fmt"
	"testing"
)

const (
	vendor_sn  = ""
	vendor_key = ""
	code       = ""
)

func TestActivate(t *testing.T) {
	ar, err := Activate(vendor_sn, vendor_key, code)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ar.BizResponse.TerminalKey, ar.BizResponse.TerminalSn)
	cir, err := CheckIn(ar.BizResponse.TerminalSn, ar.BizResponse.TerminalKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(cir)

	params := map[string]string{
		"description":  "WWWwww",
		"client_sn":    getClient_Sn(16),
		"notify_url":   "http://localhost:3000/api/payment/callback",
		"total_amount": "1",
		"return_url":   "https://www.baidu.com",
		"terminal_sn":  cir.BizResponse.TerminalSn,
		"subject":      "Pizza",
		"operator":     "kay", //门店操作员
	}
	WapApiPro(cir.BizResponse.TerminalSn, cir.BizResponse.TerminalKey, params)
}
