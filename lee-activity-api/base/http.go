package base

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"lee-activity-framework/lee-activity-api/logger"
	"net/http"
)

func HttpPost(url string, body interface{}, result interface{}) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return err
	}
	logger.Infof("http post | url: %s,req: %s,resp: %s", url, string(reqBody), string(respBody))

	return nil
}
