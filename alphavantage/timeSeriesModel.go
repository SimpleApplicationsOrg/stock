package alphavantage

import "sort"

// TimeSeriesData contains
// realtime and historical equity data in 4 different temporal resolutions:
// (1) intraday, (2) daily, (3) weekly, and (4) monthly.
// Daily, weekly, and monthly time series contain up to 20 years of historical data.
// Intraday time series typically span the past 10 to 15 active trading days.
type TimeSeriesData struct {
	MetaData   *MetaData
	TimeSeries *TimeSeries
}

// MetaData about the TimeSeriesData
type MetaData map[string]string

// Information gives the value for the metadata "1. Information"
func (m MetaData) Information() string {
	return m["1. Information"]
}

// Symbol gives the values for the metadata "2. Symbol"
func (m MetaData) Symbol() string {
	return m["2. Symbol"]
}

// LastRefreshed gives the value for the metadata "3. Last Refreshed"
func (m MetaData) LastRefreshed() string {
	return m["3. Last Refreshed"]
}

// Interval gives the value for the metadata "4. Interval"
func (m MetaData) Interval() string {
	return m["4. Interval"]
}

// OutputSize give the value for the metadata "5. Output Size"
func (m MetaData) OutputSize() string {
	return m["5. Output Size"]
}

// TimeZone give the value for the metadata "6. Time Zone"
func (m MetaData) TimeZone() string {
	return m["6. Time Zone"]
}

// TimeSeries temporal resolution with the timestamp as the keys
type TimeSeries map[string]Series

// Keys gives an array with all the value keys from the time series
func (t TimeSeries) Keys() []string {
	var timeSeriesKeys []string
	for key := range t {
		timeSeriesKeys = append(timeSeriesKeys, key)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(timeSeriesKeys)))
	return timeSeriesKeys
}

// Series is a map containing the open, high, low, close, volume values
type Series map[string]string

// Open gives the value of the "1. open"
func (s Series) Open() string {
	return s["1. open"]
}

// High gives the value of the "2. high"
func (s Series) High() string {
	return s["2. high"]
}

// Low gives the value of the "3. low"
func (s Series) Low() string {
	return s["3. low"]
}

// Close gives the value of the "4. close"
func (s Series) Close() string {
	return s["4. close"]
}

// Volume gives the value of the "5. volume"
func (s Series) Volume() string {
	return s["5. volume"]
}
