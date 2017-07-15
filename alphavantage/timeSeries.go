package alphavantage

import (
	"encoding/json"
	"fmt"
	"github.com/SimpleApplicationsOrg/stock/client"
	"github.com/tidwall/gjson"
	"log"
)

// TimeSeriesIntraday returns intraday time series (timestamp, open, high, low, close, volume) of the equity specified, updated realtime.
func (api *AVClient) TimeSeriesIntraday(symbol string, interval string) (*TimeSeriesData, error) {

	req := client.BaseAPIReq().WithPath("query").
		AddParam("function", "TIME_SERIES_INTRADAY").
		AddParam("symbol", symbol).
		AddParam("interval", interval)

	res, err := timeSeriesCall(*api.client, req)

	metaData := gjson.Get(res, "Meta Data")
	timeSeries := gjson.Get(res, fmt.Sprintf("Time Series (%s)", interval))

	metaDataIntraday, err := unmarshalMetaDataIntraday(metaData.String())
	if err != nil {
		return nil, err
	}

	timeSeriesMetaData, err := buildTimeSeriesData(metaDataIntraday, timeSeries.String())
	if err != nil {
		return nil, err
	}

	return timeSeriesMetaData, nil
}

func unmarshalMetaDataIntraday(metadata string) (*MetaData, error) {

	var timeSeriesMetaDataIntraday *MetaDataIntraday
	err := json.Unmarshal([]byte(metadata), &timeSeriesMetaDataIntraday)
	if err != nil {
		return nil, err
	}

	metaDataIntraday := MetaData(timeSeriesMetaDataIntraday)
	return &metaDataIntraday, nil
}

func timeSeriesCall(client client.APIClient, req *client.APIRequest) (string, error) {

	res, err := client.Call(req)
	if err != nil {
		log.Printf("timeSeriesCall: %s", err.Error())
		return "", err
	}

	return res, nil
}

func buildTimeSeriesData(metaData *MetaData, timeSeries string) (*TimeSeriesData, error) {

	var series TimeSeries
	err := json.Unmarshal([]byte(timeSeries), &series)
	if err != nil {
		return nil, err
	}

	return &TimeSeriesData{metaData, &series}, nil
}
