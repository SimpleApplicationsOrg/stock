package client

import "time"

const defaultTimeout time.Duration = 2 * time.Second

// Configuration of APIClient
type Configuration struct {
	url     string
	keys    map[string]string
	timeout time.Duration
}

// NewConfiguration builds a basic Configuration from a API url
func NewConfiguration(url string) *Configuration {
	return &Configuration{
		url:     url,
		keys:    make(map[string]string),
		timeout: defaultTimeout}
}

// WithTimeout sets a custom timeout in the APIClient Configuration
func (c *Configuration) WithTimeout(timeout time.Duration) *Configuration {
	c.timeout = timeout
	return c
}

// AddKey adds a key, value in the APIClient Configuration
func (c *Configuration) AddKey(key string, value string) *Configuration {
	c.keys[key] = value
	return c
}
