package comment

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//type IpS struct {
//	Origin string `json:"origin"`
//}

func GetPublicIP() (ip string, err error) {
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	Ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	json.Unmarshal(Ip, &ip)

	return ip, nil
}
