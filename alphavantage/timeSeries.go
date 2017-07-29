package alphavantage

import (
	"encoding/json"
	"fmt"
	"github.com/SimpleApplicationsOrg/stock/client"
	"github.com/tidwall/gjson"
	"log"
)

const (
	Intraday      = "TIME_SERIES_INTRADAY"
	Daily         = "TIME_SERIES_DAILY"
	DailyAdjusted = "TIME_SERIES_DAILY_ADJUSTED"
	Weekly        = "TIME_SERIES_WEEKLY"
	Monthly       = "TIME_SERIES_MONTHLY"
)

// TimeSeriesIntraday returns intraday time series (timestamp, open, high, low, close, volume) of the equity specified, updated realtime.
func (api *AVClient) TimeSeriesIntraday(symbol, interval string) (*TimeSeriesData, error) {

	req := client.NewAPIReq().WithPath("query").
		AddParam("function", Intraday).
		AddParam("symbol", symbol).
		AddParam("interval", interval)

	res, err := timeSeriesCall(*api.client, req)

	metaData := gjson.Get(res, "Meta Data")
	timeSeries := gjson.Get(res, fmt.Sprintf("Time Series (%s)", interval))

	timeSeriesMetaData, err := buildTimeSeriesData(metaData.String(), timeSeries.String())
	if err != nil {
		return nil, err
	}

	return timeSeriesMetaData, nil
}

// TimeSeries gets daily, weekly and monthly series
func (api *AVClient) TimeSeries(function, symbol string) (*TimeSeriesData, error) {

	req := client.NewAPIReq().WithPath("query").
		AddParam("function", function).
		AddParam("symbol", symbol)

	res, err := timeSeriesCall(*api.client, req)

	metaData := gjson.Get(res, "Meta Data")
	timeSeries := gjson.Get(res, timeSeriesType(function))

	timeSeriesMetaData, err := buildTimeSeriesData(metaData.String(), timeSeries.String())
	if err != nil {
		return nil, err
	}

	return timeSeriesMetaData, nil
}

func timeSeriesType(function string) string {
	var timeSeriesType string
	if function == Daily || function == DailyAdjusted {
		timeSeriesType = "Time Series (Daily)"
	} else if function == Weekly {
		timeSeriesType = "Weekly Time Series"
	} else if function == Monthly {
		timeSeriesType = "Monthly Time Series"
	}
	return timeSeriesType
}

func timeSeriesCall(client client.APIClient, req *client.APIRequest) (string, error) {

	res, err := client.Call(req)
	if err != nil {
		log.Printf("timeSeriesCall: %s", err.Error())
		return "", err
	}

	return res, nil
}

func buildTimeSeriesData(metaData, timeSeries string) (*TimeSeriesData, error) {

	var meta MetaData
	err := json.Unmarshal([]byte(metaData), &meta)
	if err != nil {
		return nil, err
	}

	var series TimeSeries
	err = json.Unmarshal([]byte(timeSeries), &series)
	if err != nil {
		return nil, err
	}

	return &TimeSeriesData{&meta, &series}, nil
}
