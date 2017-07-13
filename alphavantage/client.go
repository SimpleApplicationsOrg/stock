package alphavantage

import (
	"fmt"
	"github.com/SimpleApplicationOrg/stock/client"
	"os"
)

const (
	envURL      = "ALPHA_VANTAGE_URL"
	envKeyName  = "ALPHA_VANTAGE_KEY_NAME"
	envKeyValue = "ALPHA_VANTAGE_KEY_VALUE"
)

// Client is the specific Alpha Vantage API client
func Client() (*client.APIClient, error) {
	config, err := configure()
	if err != nil {
		return nil, err
	}
	return client.NewClient(config), nil
}

func configure() (*client.Configuration, error) {
	url := os.Getenv(envURL)
	keyName := os.Getenv(envKeyName)
	keyValue := os.Getenv(envKeyValue)

	if url == "" {
		return nil, fmt.Errorf("missing %s", envURL)
	}

	if keyName == "" {
		return nil, fmt.Errorf("missing %s", envKeyName)
	}

	if keyValue == "" {
		return nil, fmt.Errorf("missing %s", envKeyValue)
	}

	config := client.NewConfiguration(url)
	config.AddKey(keyName, keyValue)

	return config, nil
}
