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
)

func Activate(vendor_sn, vendor_key, code string) *ActivateResult {
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
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
	return nil
}

func CheckIn(terminal_sn, terminal_key string) {

}

// func WapApiPro(terminal_sn, terminal_key string, paras map[string]string) string {

// }

type ActivateResult struct {
	MerchantName string `json:"merchant_name"`
	MerchantSn   string `json:"merchant_sn"`
	StoreName    string `json:"store_name"`
	StoreSn      string `json:"store_sn"`
	TerminalKey  string `json:"terminal_key"`
	TerminalSn   string `json:"terminal_sn"`
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
