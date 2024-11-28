package sapi

import (
	"bytes"
	"crypto/tls"
	jsoniter "github.com/json-iterator/go"
	"io"
	"moul.io/http2curl"
	"net"
	"net/http"
	"time"
)

type (
	innerClient struct {
		client  *http.Client
		reqLogs []*RequestLog
	}
	RequestLog struct {
		Curl     string `json:"curl"`
		Err      error  `json:"err"`
		Status   int    `json:"status"`
		Response string `json:"response"`
	}
)

var (
	innerCacheTransport = &http.Transport{
		MaxIdleConnsPerHost:   333,
		IdleConnTimeout:       33 * time.Second,
		MaxConnsPerHost:       3000,
		ResponseHeaderTimeout: time.Second * 10,
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 10,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS11,
		},
	}

	defaultHttpTimeout = time.Minute
)

func (i *innerClient) Get(url string, params map[string]string) (resBody []byte, err error) {

	u, err := buildUrl(url, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	var (
		status int
		body   []byte
	)

	defer func() {
		i.addRequestLog(req, err, status, body)
	}()

	res, err := i.client.Do(req)
	if err != nil {
		return
	}

	status = res.StatusCode
	if res.Body != nil {
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return
		}
	}

	return body, nil

}

func (i *innerClient) Post(url string, reqParamsMapOrStruct any) (resBody []byte, err error) {

	bodyData, err := jsoniter.Marshal(reqParamsMapOrStruct)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(bodyData)

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	var (
		status int
		body   []byte
	)

	defer func() {
		reader.Seek(0, io.SeekStart)
		i.addRequestLog(req, err, status, body)
	}()

	res, err := i.client.Do(req)
	if err != nil {
		return
	}

	status = res.StatusCode
	if res.Body != nil {
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return
		}
	}

	return body, nil

}

func (i *innerClient) addRequestLog(req *http.Request, err error, status int, body []byte) {
	curl, _ := http2curl.GetCurlCommand(req)
	log := &RequestLog{
		Curl:     curl.String(),
		Err:      err,
		Status:   status,
		Response: string(body),
	}
	i.reqLogs = append(i.reqLogs, log)
}

func (i *innerClient) GetRequestLog() []*RequestLog {
	return i.reqLogs
}

func (i *innerClient) GetLastRequestLog() RequestLog {
	if len(i.reqLogs) < 0 {
		return RequestLog{}
	}
	lastLog := i.reqLogs[len(i.reqLogs)-1]
	return *lastLog
}
