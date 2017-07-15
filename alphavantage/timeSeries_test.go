package alphavantage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func response() string {
	pathResJSON := path.Join("testResources", string(os.PathSeparator), "timeSeriesIntraday.json")
	res, err := ioutil.ReadFile(pathResJSON)
	if err != nil {
		fmt.Printf("Error retrieving test response: %s", err.Error())
	}

	return string(res)
}

func testConfiguration(url string) {
	os.Setenv(envURL, url)
	os.Setenv(envKeyName, "apiKey")
	os.Setenv(envKeyValue, "testKey")
}

func buildTestServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, response())
	}))
	return ts
}

func call() (*TimeSeriesData, error) {
	ts := buildTestServer()
	defer ts.Close()
	testConfiguration(ts.URL)

	avClient, err := NewAVClient()
	if err != nil {
		return nil, err
	}

	resp, err := avClient.TimeSeriesIntraday("TEST", "1min")
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func TestMetaDataInformation(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	metaData := *resp.MetaData

	if metaData.Information() != "Intraday (1min) prices and volumes" {
		t.Errorf("metadata not properly set: Information = %s", metaData.Information())
	}
}

func TestMetaDataSymbol(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	metaData := *resp.MetaData

	if metaData.Symbol() != "TEST" {
		t.Errorf("metadata not properly set: Symbol = %s", metaData.Symbol())
	}
}

func TestMetaDataLastRefreshed(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	metaData := *resp.MetaData

	if metaData.LastRefreshed() != "2017-07-14 16:00:00" {
		t.Errorf("metadata not properly set: LastRefreshed = %s", metaData.LastRefreshed())
	}
}

func TestMetaDataInterval(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	metaData := *resp.MetaData

	if metaData.Interval() != "1min" {
		t.Errorf("metadata not properly set: Interval = %s", metaData.Interval())
	}
}

func TestMetaDataOutputSize(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	metaData := *resp.MetaData

	if metaData.OutputSize() != "Compact" {
		t.Errorf("metadata not properly set: OutputSize = %s", metaData.OutputSize())
	}
}

func TestMetaDataTimeZone(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	metaData := *resp.MetaData

	if metaData.TimeZone() != "US/Eastern" {
		t.Errorf("metadata not properly set: TimeZone = %s", metaData.TimeZone())
	}
}

func TestTimeSeriesKeys(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	timeSeries := *resp.TimeSeries
	count := len(timeSeries.Keys())
	if count != 100 {
		t.Errorf("timeSeries not properly set: keys count = %d", count)
	}
}

func TestTimeSeriesOpen(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	timeSeries := *resp.TimeSeries
	series := timeSeries["2017-07-14 16:00:00"]

	if series.Open() != "72.8900" {
		t.Errorf("timeSeries not properly set: Open = %s", series.Open())
	}
}

func TestTimeSeriesHigh(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	timeSeries := *resp.TimeSeries
	series := timeSeries["2017-07-14 16:00:00"]

	if series.High() != "72.9000" {
		t.Errorf("timeSeries not properly set: High = %s", series.High())
	}
}

func TestTimeSeriesLow(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	timeSeries := *resp.TimeSeries
	series := timeSeries["2017-07-14 16:00:00"]

	if series.Low() != "72.7500" {
		t.Errorf("timeSeries not properly set: High = %s", series.Low())
	}
}

func TestTimeSeriesClose(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	timeSeries := *resp.TimeSeries
	series := timeSeries["2017-07-14 16:00:00"]

	if series.Close() != "72.7800" {
		t.Errorf("timeSeries not properly set: High = %s", series.Close())
	}
}

func TestTimeSeriesVolume(t *testing.T) {
	resp, err := call()
	if err != nil {
		t.Errorf("error on call: %s", err.Error())
	}

	timeSeries := *resp.TimeSeries
	series := timeSeries["2017-07-14 16:00:00"]

	if series.Volume() != "2617441" {
		t.Errorf("timeSeries not properly set: High = %s", series.Volume())
	}
}
