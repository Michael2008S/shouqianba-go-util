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

const (
	TerminalSn  = ""
	TerminalKey = ""
)

func TestActivate(t *testing.T) {
	// ar, err := Activate(vendor_sn, vendor_key, code)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(ar.BizResponse.TerminalKey, ar.BizResponse.TerminalSn)

	params := map[string]string{
		"description":  "WWWwww",
		"client_sn":    getClient_Sn(16),
		"notify_url":   "http://www.baidu.com",
		"total_amount": "10",
		"return_url":   "https://www.baidu.com",
		"terminal_sn":  TerminalSn,
		"subject":      "Pizza",
		"operator":     "kay", //门店操作员
	}
	payurl := WapApiPro(TerminalSn, TerminalKey, params)

	fmt.Println("payurl++++>", payurl)
}
