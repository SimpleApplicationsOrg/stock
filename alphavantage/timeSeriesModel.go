package alphavantage

import "sort"

type TimeSeriesData struct {
	MetaData   *MetaData
	TimeSeries *TimeSeries
}

type MetaData map[string]string

func (m MetaData) Information() string {
	return m["1. Information"]
}

func (m MetaData) Symbol() string {
	return m["2. Symbol"]
}

func (m MetaData) LastRefreshed() string {
	return m["3. Last Refreshed"]
}

func (m MetaData) Interval() string {
	return m["4. Interval"]
}

func (m MetaData) OutputSize() string {
	return m["5. Output Size"]
}

func (m MetaData) TimeZone() string {
	return m["6. Time Zone"]
}

type TimeSeries map[string]Series

func (t TimeSeries) Keys() []string {
	var timeSeriesKeys []string
	for key := range t {
		timeSeriesKeys = append(timeSeriesKeys, key)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(timeSeriesKeys)))
	return timeSeriesKeys
}

type Series map[string]string

func (s Series) Open() string {
	return s["1. open"]
}

func (s Series) High() string {
	return s["2. high"]
}

func (s Series) Low() string {
	return s["3. low"]
}

func (s Series) Close() string {
	return s["4. close"]
}

func (s Series) Volume() string {
	return s["5. volume"]
}
