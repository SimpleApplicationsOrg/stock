package client

import "time"

// Configuration of APIClient
type Configuration struct {
	URL     string
	Keys    map[string]string
	Timeout time.Duration
}

// DefaultTimeout for a APIClient Configuration
const DefaultTimeout time.Duration = 2 * time.Second

// NewConfiguration builds a basic Configuration from a API URL
func NewConfiguration(url string) *Configuration {
	return &Configuration{
		URL:     url,
		Keys:    make(map[string]string),
		Timeout: DefaultTimeout}
}

// WithTimeout sets a custom timeout in the APIClient Configuration
func (c *Configuration) WithTimeout(timeout time.Duration) *Configuration {
	c.Timeout = timeout
	return c
}

// AddKey adds a key, value in the APIClient Configuration
func (c *Configuration) AddKey(key string, value string) *Configuration {
	c.Keys[key] = value
	return c
}
