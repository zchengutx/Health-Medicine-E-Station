package comment

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IpResponse struct {
	Origin string `json:"origin"`
}

func GetPublicIP() (ip string, err error) {
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ipResp IpResponse
	err = json.Unmarshal(body, &ipResp)
	if err != nil {
		return "", err
	}

	return ipResp.Origin, nil
}
