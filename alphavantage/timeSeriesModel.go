package alphavantage

import (
	"sort"
	"strings"
)

// TimeSeriesData contains
// realtime and historical equity data in 4 different temporal resolutions:
// (1) intraday, (2) daily, (3) weekly, and (4) monthly.
// Daily, weekly, and monthly time series contain up to 20 years of historical data.
// Intraday time series typically span the past 10 to 15 active trading days.
type TimeSeriesData struct {
	MetaData   *MetaData
	TimeSeries *TimeSeries
}

const (
	information   = "Information"
	symbol        = "Symbol"
	lastRefreshed = "Last Refreshed"
	interval      = "Interval"
	outputSize    = "Output Size"
	timeZone      = "Time Zone"
)

// MetaData about the time series data
type MetaData map[string]string

// Information gives the metadata value
func (m MetaData) Information() string {
	return m[m.key(information)]
}

// Symbol gives the values for that metadata
func (m MetaData) Symbol() string {
	return m[m.key(symbol)]
}

// LastRefreshed gives the value for that metadata
func (m MetaData) LastRefreshed() string {
	return m[m.key(lastRefreshed)]
}

// Interval gives the value for that metadata
func (m MetaData) Interval() string {
	return m[m.key(interval)]
}

// OutputSize give the value for that metadata
func (m MetaData) OutputSize() string {
	return m[m.key(outputSize)]
}

// TimeZone give the value for that metadata
func (m MetaData) TimeZone() string {
	return m[m.key(timeZone)]
}

func (m MetaData) key(contains string) string {
	for key := range m {
		if strings.Contains(key, contains) {
			return key
		}
	}
	return ""
}

// TimeSeries temporal resolution with the timestamp as the keys
type TimeSeries map[string]Series

// TimeStamps gives an array with all the value keys from the time series
func (t TimeSeries) TimeStamps() []string {
	var timeStamps []string
	for key := range t {
		timeStamps = append(timeStamps, key)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(timeStamps)))
	return timeStamps
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
