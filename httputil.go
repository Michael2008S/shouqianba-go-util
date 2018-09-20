package shouqianba

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_DOMAIN = "https://api.shouqianba.com"

	ACTIVATE_URL = "/terminal/activate"
	CHECKIN_URL  = "/terminal/checkin"

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

func WapApiPro(terminal_sn, terminal_key string, params map[string]string) string {
	sortMap(params)

	return WAP_API_PRO_URL
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
