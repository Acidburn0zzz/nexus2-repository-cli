package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(url, method string, body []byte, username, password string, verbose bool) ([]byte, string) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	//Print verbose logs if verbose flag is set
	if verbose {
		fmt.Println("Request Url:", req.URL)
		fmt.Println("Request Headers:", req.Header)
		fmt.Println("Response Headers:", resp.Header)
		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Body:", string(responseBody))
	}

	return responseBody, resp.Status
}
