package netatmo

import (
	"net/url"
	"strconv"
)

type DashboardData struct {
	AbsolutePressure *float64 `json:"AbsolutePressure,omitempty"`
	CO2              *int64   `json:"CO2,omitempty"`
	Humidity         *int64   `json:"Humidity,omitempty"`
	Noise            *int64   `json:"Noise,omitempty"`
	Pressure         *float64 `json:"Pressure,omitempty"`
	Temperature      *float64 `json:"Temperature,omitempty"`
	DateMaxTemp      *int64   `json:"date_max_temp,omitempty"`
	DateMinTemp      *int64   `json:"date_min_temp,omitempty"`
	MaxTemp          *float64 `json:"max_temp,omitempty"`
	MinTemp          *float64 `json:"min_temp,omitempty"`
	PressureTrend    *string  `json:"pressure_trend,omitempty"`
	TempTrend        *string  `json:"temp_trend,omitempty"`
	TimeUTC          *int64   `json:"time_utc,omitempty"`
}

type Module struct {
	// ID             *string `json:"_id,omitempty"`
	// LastMessage    *int64  `json:"last_message,omitempty"`
	// LastSeen       *int64  `json:"last_seen,omitempty"`
	// BatteryVP      *int64  `json:"battery_vp,omitempty"`
	// BatteryPercent *int64  `json:"battery_percent,omitempty"`
	// RfStatus       *int64  `json:"rf_status,omitempty"`
}

type Place struct {
	City     *string   `json:"city,omitempty"`
	Country  *string   `json:"country,omitempty"`
	Location []float64 `json:"location,omitempty"`
	Timezone *string   `json:"timezone,omitempty"`
}

type Device struct {
	ID              *string        `json:"_id,omitempty"`
	CipherID        *string        `json:"cipher_id,omitempty"`
	CO2Calibrating  *bool          `json:"co2_calibrating,omitempty"`
	DashboardData   *DashboardData `json:"dashboard_data,omitempty"`
	DataType        []string       `json:"data_type,omitempty"` // Temperature, CO2, Humidity, Noise, Pressure
	DataSetup       *int64         `json:"data_setup,omitempty"`
	Firmware        *int64         `json:"firmware,omitempty"`
	LastSetup       *int64         `json:"last_setup,omitempty"`
	LastStatusStore *int64         `json:"last_status_store,omitempty"`
	LastUpgrade     *int64         `json:"last_upgrade,omitempty"`
	ModuleName      *string        `json:"module_name,omitempty"`
	Modules         []*Module      `json:"modules,omitempty"`
	Place           *Place         `json:"place,omitempty"`
	StationName     *string        `json:"station_name,omitempty"`
	Type            *string        `json:"type,omitempty"`
	WifiStatus      *int64         `json:"wifi_status,omitempty"`
}

type Administrative struct {
	FeelLikeAlgo *int64  `json:"feel_like_algo,omitempty"`
	Lang         *string `json:"lang,omitempty"`
	Pressureunit *int64  `json:"pressureunit,omitempty"`
	RegLocale    *string `json:"reg_locale,omitempty"`
	Unit         *int64  `json:"unit,omitempty"`
	Windunit     *int64  `json:"windunit,omitempty"`
}

type User struct {
	Administrative *Administrative `json:"administrative,omitempty"`
	Mail           *string         `json:"mail,omitempty"`
}

type Body struct {
	Devices []*Device `json:"devices,omitempty"`
	User    *User     `json:"user,omitempty"`
}

type StationData struct {
	Body       *Body    `json:"body,omitempty"`
	Status     *string  `json:"status,omitempty"`
	TimeExec   *float64 `json:"time_exec,omitempty"`
	TimeServer *int64   `json:"time_server,omitempty"`
}

// https://dev.netatmo.com/en-US/resources/technical/reference/weatherstation/getstationsdata
// Returns data from a user Weather Stations (measures and device specific data)
// scope: read_station
func (c *Client) Getstationsdata(deviceID string, getFavorites bool) (*StationData, error) {
	u, err := url.Parse("https://api.netatmo.com/api/getstationsdata")
	if err != nil {
		return nil, err
	}

	queries := url.Values{}
	queries.Add("access_token", c.token)
	if len(deviceID) > 0 {
		queries.Add("device_id", deviceID)
	}
	queries.Add("get_favorites", strconv.FormatBool(getFavorites))

	u.RawQuery = queries.Encode()

	var d StationData
	if err := c.get(u, &d); err != nil {
		return nil, err
	}

	return &d, nil
}
