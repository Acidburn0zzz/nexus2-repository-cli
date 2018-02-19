package utils

import (
	"bytes"
	m "com/privateSquare/go/nexus-repository-cli/model"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
CreateBaseRequest create the base request for a HTTP request
@param method  	string        	http request method eg: GET, POST, etc
@param url     	string    		http request url
@param body     []byte    		request body
@param user    	m.AuthUser  	User authentication details
@param verbose  boolean    		prints verbose logs if set to true
@return *http.Request	HTTP base request
*/
func CreateBaseRequest(method, url string, body []byte, user m.AuthUser, verbose bool) *http.Request {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.SetBasicAuth(user.Username, user.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	Error(err, "Error creating the request")

	if verbose {
		fmt.Println("Request Url:", req.URL)
		fmt.Println("Request Headers:", req.Header)
		fmt.Println("Request Body:", req.Body)
	}

	return req
}

/*
HTTPRequest makes a request to the remote server
@param user		m.AuthUser		User authentication details
@param req		*http.Request   HTTP base request
@param verbose  boolean    		prints verbose logs if set to true
@return []byte	response body
@return string	response status
*/
func HTTPRequest(user m.AuthUser, req *http.Request, verbose bool) ([]byte, string) {

	client := &http.Client{}
	resp, err := client.Do(req)
	Error(err, "There was a problem in making the request")

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	Error(err, "There was a problem reading the response body")

	//b.CheckAuth(user, resp.Status)

	if verbose {
		fmt.Println("Response Headers:", resp.Header)
		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Body:", string(respBody))
	}
	return respBody, resp.Status
}