package sapi

import (
	"net/http"
	"time"
)

func NewClient() *innerClient {
	return &innerClient{client: &http.Client{
		Transport: innerCacheTransport,
		Timeout:   defaultHttpTimeout,
	},
		reqLogs: make([]*RequestLog, 0)}
}

func NewClientWithTimeout(timeout time.Duration) *innerClient {
	return &innerClient{client: &http.Client{
		Transport: innerCacheTransport,
		Timeout:   timeout,
	},
		reqLogs: make([]*RequestLog, 0),
	}
}
