package network

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type (
	HttpClientRequest struct {
		Method       string
		Url          string
		InsecureSkip bool
		Headers      map[string]string
		Params       map[string]string
		Data         interface{}
	}

	HttpClientResponse struct {
		Status        int
		StatusMsg     string
		IsStatusIn2xx bool
		Url           string
		Headers       map[string]string
		Request       HttpClientRequest
		Data          map[string]interface{}
		Error         error
	}

	HttpClientResponseRaw struct {
		Status        int
		StatusMsg     string
		IsStatusIn2xx bool
		Url           string
		Headers       map[string]string
		Request       HttpClientRequest
		Data          string
		Error         error
	}
)

type HttpClient struct {
}

func (obj HttpClient) Call(data HttpClientRequest) HttpClientResponse {
	var body []byte
	var request *http.Request
	var err error

	//Adding TLS Security - remove keep alive to avoid multi connection
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DisableKeepAlives = true
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: data.InsecureSkip,
	}

	client := &http.Client{
		Transport: tr,
	}

	// Adding body
	if data.Data != nil {
		body, err = json.Marshal(data.Data)

		if err != nil {
			return HttpClientResponse{Error: err, Request: data}
		}

		request, err = http.NewRequest(data.Method, data.Url, bytes.NewBuffer(body))

		if request.Body != nil {
			defer request.Body.Close()
		}
	} else {
		request, err = http.NewRequest(data.Method, data.Url, nil)
	}

	if err != nil {
		return HttpClientResponse{Error: err, Request: data}
	}

	//Adding Headers
	for key, value := range data.Headers {
		request.Header.Add(key, value)
	}

	request.Header.Add("Content-Type", "application/json")

	//Adding Parameters
	query := request.URL.Query()
	for key, value := range data.Params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()

	//Execute request
	finalurl := request.URL.String()
	result, err := client.Do(request)

	if err != nil {
		return HttpClientResponse{Error: err, Request: data, Url: finalurl}
	}

	defer result.Body.Close()
	body, err = ioutil.ReadAll(result.Body)

	if err != nil {
		return HttpClientResponse{Error: err, Request: data, Url: finalurl}
	}

	//Build response
	response := HttpClientResponse{
		Error:   nil,
		Request: data,
		Headers: map[string]string{},
		Url:     finalurl,
	}

	//Getting body in JSON
	err = json.Unmarshal(body, &response.Data)

	if err != nil {
		response.Error = err
		return response
	}

	//Status Code
	response.IsStatusIn2xx = result.StatusCode >= 200 && result.StatusCode < 300
	response.Status = result.StatusCode
	response.StatusMsg = result.Status

	//Getting headers
	for key := range result.Header {
		response.Headers[key] = result.Header.Get(key)
	}

	return response
}

func (obj HttpClient) CallForm(data HttpClientRequest) HttpClientResponse {
	var body []byte
	var request *http.Request
	var err error

	//Adding TLS Security - remove keep alive to avoid multi connection
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DisableKeepAlives = true
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: data.InsecureSkip,
	}

	client := &http.Client{
		Transport: tr,
	}

	values := url.Values{}
	//Adding Parameters
	for key, value := range data.Params {
		values.Add(key, value)
	}

	request, err = http.NewRequest(data.Method, data.Url, strings.NewReader(values.Encode()))

	if err != nil {
		return HttpClientResponse{Error: err, Request: data}
	}

	//Adding Headers
	for key, value := range data.Headers {
		request.Header.Add(key, value)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//Execute request
	finalurl := request.URL.String()
	result, err := client.Do(request)

	if err != nil {
		return HttpClientResponse{Error: err, Request: data, Url: finalurl}
	}

	defer result.Body.Close()
	body, err = ioutil.ReadAll(result.Body)

	if err != nil {
		return HttpClientResponse{Error: err, Request: data, Url: finalurl}
	}

	//Build response
	response := HttpClientResponse{
		Error:   nil,
		Request: data,
		Headers: map[string]string{},
		Url:     finalurl,
	}

	//Getting body in JSON
	err = json.Unmarshal(body, &response.Data)

	if err != nil {
		response.Error = err
		return response
	}

	//Status Code
	response.IsStatusIn2xx = result.StatusCode >= 200 && result.StatusCode < 300
	response.Status = result.StatusCode
	response.StatusMsg = result.Status

	//Getting headers
	for key := range result.Header {
		response.Headers[key] = result.Header.Get(key)
	}

	return response
}

func (obj HttpClient) CallFormRaw(data HttpClientRequest) HttpClientResponseRaw {
	var body []byte
	var request *http.Request
	var err error

	//Adding TLS Security - remove keep alive to avoid multi connection
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DisableKeepAlives = true
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: data.InsecureSkip,
	}

	client := &http.Client{
		Transport: tr,
	}

	values := url.Values{}
	//Adding Parameters
	for key, value := range data.Params {
		values.Add(key, value)
	}

	request, err = http.NewRequest(data.Method, data.Url, strings.NewReader(values.Encode()))

	if err != nil {
		return HttpClientResponseRaw{Error: err, Request: data}
	}

	//Adding Headers
	for key, value := range data.Headers {
		request.Header.Add(key, value)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//Execute request
	finalurl := request.URL.String()
	result, err := client.Do(request)

	if err != nil {
		return HttpClientResponseRaw{Error: err, Request: data, Url: finalurl}
	}

	defer result.Body.Close()
	body, err = ioutil.ReadAll(result.Body)

	if err != nil {
		return HttpClientResponseRaw{Error: err, Request: data, Url: finalurl}
	}

	//Build response
	response := HttpClientResponseRaw{
		Error:   nil,
		Request: data,
		Headers: map[string]string{},
		Url:     finalurl,
	}

	//Getting body in RAW
	response.Data = string(body)

	if err != nil {
		response.Error = err
		return response
	}

	//Status Code
	response.IsStatusIn2xx = result.StatusCode >= 200 && result.StatusCode < 300
	response.Status = result.StatusCode
	response.StatusMsg = result.Status

	//Getting headers
	for key := range result.Header {
		response.Headers[key] = result.Header.Get(key)
	}

	return response
}
