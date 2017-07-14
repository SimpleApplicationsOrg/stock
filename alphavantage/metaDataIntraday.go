package alphavantage

// MetaDataIntraday implementation
type MetaDataIntraday map[string]string

// Information gives the value for the metadata "1. Information"
func (m MetaDataIntraday) Information() string {
	return m["1. Information"]
}

// Symbol gives the values for the metadata "2. Symbol"
func (m MetaDataIntraday) Symbol() string {
	return m["2. Symbol"]
}

// LastRefreshed gives the value for the metadata "3. Last Refreshed"
func (m MetaDataIntraday) LastRefreshed() string {
	return m["3. Last Refreshed"]
}

// Interval gives the value for the metadata "4. Interval"
func (m MetaDataIntraday) Interval() string {
	return m["4. Interval"]
}

// OutputSize give the value for the metadata "5. Output Size"
func (m MetaDataIntraday) OutputSize() string {
	return m["5. Output Size"]
}

// TimeZone give the value for the metadata "6. Time Zone"
func (m MetaDataIntraday) TimeZone() string {
	return m["6. Time Zone"]
}
