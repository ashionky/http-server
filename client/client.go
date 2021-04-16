/**
 * @Author pibing
 */

package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)


type CliData struct {
	ErrMsg string  `json:"err_msg"`
	Barr   []bool  `json:"bool_arr"`
}

//BcjClient方法
func BcjClient(arr []string) ([]bool,error) {
	url:="https://127.0.0.1:9999/api"
	data :=map[string]interface{}{
		"req_arr":arr,
	}
	contentType:="application/json"
	s, err := Post(url, data, contentType)
	if err != nil{
		return nil,err
	}
	var cliData CliData
	err =json.Unmarshal([]byte(s),&cliData)
	if err != nil {
		return nil,err
	}
	if cliData.ErrMsg != "" {
		return cliData.Barr,errors.New(cliData.ErrMsg)
	}
	return cliData.Barr,nil
}



func Post(url string, data interface{}, contentType string) (string,error) {

	//https配置
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "",err
	}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	return string(result),nil
}
