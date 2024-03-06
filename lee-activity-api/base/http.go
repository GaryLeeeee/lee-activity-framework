package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpPost(url string, body interface{}, result interface{}) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	fmt.Printf("http post | url:%+v,body:%+v,resp:%+v", url, reqBody, resp)
	if err != nil {
		return err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(respBody, &result)
}
