package sip

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Transport: &http.Transport{
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
		},
		Timeout: 30 * time.Second,
	}
}

func Init(httpClient *http.Client) {
	client = httpClient
}

func NewClient() *innerClient {
	return &innerClient{client: client}
}

func NewClientWithTimeout(timeout time.Duration) *innerClient {
	return &innerClient{client: client, timeout: timeout}
}
