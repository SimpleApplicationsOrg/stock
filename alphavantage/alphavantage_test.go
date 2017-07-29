package alphavantage_test

import (
	"fmt"
	"github.com/SimpleApplicationsOrg/stock/alphavantage"
)

func Example() {
	avClient, err := alphavantage.NewAVClient()
	if err != nil {
		fmt.Printf("error getting client: %s", err.Error())
		return
	}

	response, err := avClient.TimeSeriesIntraday("GOOGL", "1min")
	if err != nil {
		fmt.Printf("error calling api: %s", err.Error())
		return
	}

	metaData := *response.MetaData
	fmt.Println(metaData.Information(), metaData.OutputSize())
	fmt.Println(metaData.LastRefreshed(), metaData.TimeZone())
	fmt.Println(metaData.Interval())

	timeSeries := *response.TimeSeries
	for _, timeStamp := range timeSeries.TimeStamps() {
		value := (timeSeries)[timeStamp]
		fmt.Println(timeStamp, value.Open(), value.High(), value.Low(), value.Close(), value.Volume())
	}
}
