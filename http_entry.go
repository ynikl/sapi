package sip

func Get(url string, params map[string]string) ([]byte, string, error) {
	c := NewClient()
	body, err := c.Get(url, params)
	log := c.GetLastRequestLog()
	return body, log.Curl, err
}

func Post(url string, reqParamsMapOrStruct any) (data []byte, curl string, err error) {
	c := NewClient()
	body, err := c.Post(url, reqParamsMapOrStruct)
	log := c.GetLastRequestLog()
	return body, log.Curl, err
}
