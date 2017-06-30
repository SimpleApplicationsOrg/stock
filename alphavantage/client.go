package alphavantage

import (
	"fmt"
	"github.com/SimpleApplicationOrg/stock/client"
	"os"
)

const (
	ALPHA_VANTAGE_URL       = "ALPHA_VANTAGE_URL"
	ALPHA_VANTAGE_KEY_NAME  = "ALPHA_VANTAGE_KEY_NAME"
	ALPHA_VANTAGE_KEY_VALUE = "ALPHA_VANTAGE_KEY_VALUE"
)

func Client() (*client.APIClient, error) {
	config, err := configure()
	if err != nil {
		return nil, err
	}
	return client.NewClient(config), nil
}

func configure() (*client.Configuration, error) {
	url := os.Getenv(ALPHA_VANTAGE_URL)
	keyName := os.Getenv(ALPHA_VANTAGE_KEY_NAME)
	keyValue := os.Getenv(ALPHA_VANTAGE_KEY_VALUE)

	if url == "" {
		return nil, fmt.Errorf("missing %s", ALPHA_VANTAGE_URL)
	}

	if keyName == "" {
		return nil, fmt.Errorf("missing %s", ALPHA_VANTAGE_KEY_NAME)
	}

	if keyValue == "" {
		return nil, fmt.Errorf("missing %s", ALPHA_VANTAGE_KEY_VALUE)
	}

	config := client.NewConfiguration(url)
	config.WithKey(keyName, keyValue)

	return config, nil
}
