package network

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func DoHttpRead(client *http.Client, request *http.Request) (string, int) {
	resp, err := client.Do(request)
	if err == nil {
		ce := resp.Header.Get("Content-Encoding")
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		//var buff bytes.Buffer
		//io.Copy(&buff, resp.Body)
		//body := buff.Bytes()
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return err.Error(), 500
		}
		if ce == "gzip" {
			buf := bytes.NewBuffer(body)
			gr, _ := gzip.NewReader(buf)
			defer gr.Close()
			body, _ = ioutil.ReadAll(gr)
		}
		return string(body), resp.StatusCode
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	return err.Error(), 500
}

func GetTo(full, bearer string) (string, int) {
	fmt.Println("here")
	request, err := http.NewRequest("GET", full, nil)
	if err != nil {
		return "bad url", 500
	}
	fmt.Println("here2")
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 150}

	fmt.Println("here3")
	return DoHttpRead(client, request)
}

func GetToWith(full, name, bearer string) (string, int) {
	request, err := http.NewRequest("GET", full, nil)
	if err != nil {
		return "bad url", 500
	}
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", name, bearer))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 150}

	return DoHttpRead(client, request)
}

func PostTo(full, bearer string, payload map[string]any) (string, int) {
	asBytes, _ := json.Marshal(payload)
	body := bytes.NewBuffer(asBytes)
	request, err := http.NewRequest("POST", full, body)
	if err != nil {
		return "bad url", 500
	}
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 150}

	return DoHttpRead(client, request)
}

func PostToForm(full, bearer string, payload map[string]any) (string, int) {
	data := url.Values{}
	for key, value := range payload {
		data.Set(key, fmt.Sprintf("%v", value))
	}

	body := strings.NewReader(data.Encode())

	request, err := http.NewRequest("POST", full, body)
	if err != nil {
		return "bad url", 500
	}
	SetHeaders(bearer, request)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{Timeout: time.Second * 150}

	return DoHttpRead(client, request)
}

func SetHeaders(bearer string, request *http.Request) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearer))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Accept-Encoding", "gzip")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Max-Keep-Alive-Requests", "100")
}
