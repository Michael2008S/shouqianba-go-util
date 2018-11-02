package shouqianba

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	API_DOMAIN = "https://api.shouqianba.com"

	ACTIVATE_URL = "/terminal/activate"
	CHECKIN_URL  = "/terminal/checkin"

	QUERY_URL = "/upay/v2/query"

	CANCEL_URL = "/upay/v2/cancel"

	REVOKE_URL = "/upay/v2/revoke"

	WAP_API_PRO_URL = "https://m.wosai.cn/qr/gateway?"
)

func Activate(vendor_sn, vendor_key, code string) (ActivateResult, error) {
	var ar ActivateResult
	m := map[string]string{
		"code":        code,
		"type":        "2",
		"os_info":     "Mac OS",
		"device_id":   "50a87771-ca8a-4952-a493-9504c39ab495",
		"sdk_version": "Java SDK v1.0",
	}
	mJson, _ := json.Marshal(m)

	sign := getSign(string(mJson) + vendor_key)
	body, err := httpPost(API_DOMAIN+ACTIVATE_URL, mJson, sign, vendor_sn)
	if err != nil {
		return ar, err
	}
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &ar); err != nil {
		return ar, err
	}
	return ar, nil
}

func CheckIn(terminal_sn, terminal_key string) (CheckInResult, error) {
	var cir CheckInResult
	m := map[string]string{
		"terminal_sn": terminal_sn,
		"os_info":     "Mac OS",
		"device_id":   "50a87771-ca8a-4952-a493-9504c39ab495",
		"sdk_version": "Java SDK v1.0",
	}
	mJson, _ := json.Marshal(m)
	sign := getSign(string(mJson) + terminal_key)
	body, err := httpPost(API_DOMAIN+CHECKIN_URL, mJson, sign, terminal_sn)
	if err != nil {
		return cir, err
	}
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &cir); err != nil {
		return cir, err
	}
	return cir, nil
}

func Query(terminal_sn, terminal_key, sn, client_sn string) (QueryResult, error) {
	var query QueryResult
	m := map[string]string{
		"terminal_sn": terminal_sn,
		"sn":          sn,
		"client_sn":   client_sn,
	}
	mJson, _ := json.Marshal(m)
	sign := getSign(string(mJson) + terminal_key)
	body, err := httpPost(API_DOMAIN+QUERY_URL, mJson, sign, terminal_sn)
	if err != nil {
		return query, err
	}
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &query); err != nil {
		return query, err
	}
	return query, nil
}

//自动撤单
func Cancel(terminal_sn, terminal_key, sn, client_sn string) {
	m := map[string]string{
		"terminal_sn": terminal_sn,
		"sn":          sn,
		"client_sn":   client_sn,
	}
	mJson, _ := json.Marshal(m)
	sign := getSign(string(mJson) + terminal_key)
	body, err := httpPost(API_DOMAIN+CANCEL_URL, mJson, sign, terminal_sn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---cancel--->", string(body))
}

//手动撤单 (退款)
func Revoke(terminal_sn, terminal_key, sn, client_sn string) {
	m := map[string]string{
		"terminal_sn": terminal_sn,
		"sn":          sn,
		"client_sn":   client_sn,
	}
	mJson, _ := json.Marshal(m)
	sign := getSign(string(mJson) + terminal_key)
	body, err := httpPost(API_DOMAIN+REVOKE_URL, mJson, sign, terminal_sn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("----revoke--->", string(body))
}

func WapApiPro(terminal_sn, terminal_key string, params map[string]string) string {
	sortStr := sortMap(params)
	fmt.Println(sortStr)
	sign := strings.ToUpper(getSign(sortStr + "&key=" + terminal_key))
	return WAP_API_PRO_URL + sortStr + "&sign=" + sign
}

type ActivateResult struct {
	BizResponse struct {
		MerchantName string `json:"merchant_name"`
		MerchantSn   string `json:"merchant_sn"`
		StoreName    string `json:"store_name"`
		StoreSn      string `json:"store_sn"`
		TerminalKey  string `json:"terminal_key"`
		TerminalSn   string `json:"terminal_sn"`
	} `json:"biz_response"`
	ResultCode string `json:"result_code"`
}

type CheckInResult struct {
	BizResponse struct {
		MerchantName string `json:"merchant_name"`
		MerchantSn   string `json:"merchant_sn"`
		StoreName    string `json:"store_name"`
		StoreSn      string `json:"store_sn"`
		TerminalKey  string `json:"terminal_key"`
		TerminalSn   string `json:"terminal_sn"`
	} `json:"biz_response"`
	ResultCode string `json:"result_code"`
}

type QueryResult struct {
	BizResponse struct {
		Data struct {
			ChannelFinishTime string `json:"channel_finish_time"`
			ClientSn          string `json:"client_sn"`
			ClientTsn         string `json:"client_tsn"`
			Description       string `json:"description"`
			FinishTime        string `json:"finish_time"`
			NetAmount         string `json:"net_amount"`
			Operator          string `json:"operator"`
			OrderStatus       string `json:"order_status"`
			PayerLogin        string `json:"payer_login"`
			PayerUID          string `json:"payer_uid"`
			PaymentList       []struct {
				AmountTotal string `json:"amount_total"`
				Type        string `json:"type"`
			} `json:"payment_list"`
			Payway      string `json:"payway"`
			PaywayName  string `json:"payway_name"`
			Sn          string `json:"sn"`
			Status      string `json:"status"`
			SubPayway   string `json:"sub_payway"`
			Subject     string `json:"subject"`
			TotalAmount string `json:"total_amount"`
			TradeNo     string `json:"trade_no"`
		} `json:"data"`
		ErrorCode    string `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		ResultCode   string `json:"result_code"`
	} `json:"biz_response"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	ResultCode   string `json:"result_code"`
}

func httpPost(url string, mJson []byte, sign, sn string) ([]byte, error) {
	contentReader := bytes.NewReader(mJson)
	req, _ := http.NewRequest("POST", url, contentReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", sn+" "+sign)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
