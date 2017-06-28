package client

import "time"

type Configuration struct {
	URL     string
	Keys    map[string]string
	Timeout time.Duration
}

const DefaultTimeout time.Duration = 2 * time.Second

func NewConfiguration(url string) Configuration {
	return Configuration{
		URL:     url,
		Keys:    make(map[string]string),
		Timeout: DefaultTimeout}
}

func (c Configuration) WithTimeout(timeout time.Duration) Configuration {
	c.Timeout = timeout
	return c
}

func (c Configuration) WithKey(key string, value string) Configuration {
	c.Keys[key] = value
	return c
}
