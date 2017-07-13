package alphavantage

import (
	"encoding/json"
	"fmt"
	"github.com/SimpleApplicationOrg/stock/client"
	"github.com/tidwall/gjson"
	"log"
)

// TimeSeriesIntraday returns intraday time series (timestamp, open, high, low, close, volume) of the equity specified, updated realtime.
func TimeSeriesIntraday(apiClient *client.APIClient, symbol string, interval string) (*TimeSeriesData, error) {
	req := client.BaseAPIReq().WithPath("query").
		AddParam("function", "TIME_SERIES_INTRADAY").
		AddParam("symbol", symbol).
		AddParam("interval", interval)

	res, err := apiClient.Call(req)
	if err != nil {
		log.Printf("timeSeriesIntraday: %s", err.Error())
		return nil, err
	}

	metaData := gjson.Get(res, "Meta Data")
	timeSeries := gjson.Get(res, fmt.Sprintf("Time Series (%s)", interval))

	var timeSeriesMetaData MetaData
	err = json.Unmarshal([]byte(metaData.String()), &timeSeriesMetaData)
	if err != nil {
		return nil, err
	}

	var series TimeSeries
	err = json.Unmarshal([]byte(timeSeries.String()), &series)
	if err != nil {
		return nil, err
	}

	return &TimeSeriesData{&timeSeriesMetaData, &series}, nil
}
