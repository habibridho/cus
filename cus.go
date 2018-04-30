package cus

import (
"net/http"
"bytes"
"encoding/json"
"io/ioutil"
)

func Get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Post(url string, body interface{}) ([]byte, error) {
	bytesBuffer := new(bytes.Buffer)
	json.NewEncoder(bytesBuffer).Encode(body)
	response, err := http.Post(url, "application/json", bytesBuffer)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func Put(url string, body interface{}) ([]byte, error) {
	bytesBuffer := new(bytes.Buffer)
	json.NewEncoder(bytesBuffer).Encode(body)
	request, err := http.NewRequest(http.MethodPut, url, bytesBuffer)
	request.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}
	response, err := http.Client{}.Do(request)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
