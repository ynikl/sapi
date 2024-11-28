package sapi

import (
	urlPkg "net/url"
)

func buildUrl(url string, params map[string]string) (*urlPkg.URL, error) {

	u, err := urlPkg.Parse(url)
	if err != nil {
		return u, err
	}

	queryParams := urlPkg.Values{}
	for k, v := range params {
		queryParams.Add(k, v)
	}
	u.RawQuery = queryParams.Encode()
	return u, nil
}
