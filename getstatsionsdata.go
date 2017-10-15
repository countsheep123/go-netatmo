package netatmo

import (
	"net/url"
	"strconv"
)

type DashboardData struct {
	TimeUTC     *int64   `json:"time_utc,omitempty"`
	Temperature *float64 `json:"Temperature,omitempty"`
	TempTrend   *string  `json:"temp_trend,omitempty"`
	Humidity    *int64   `json:"Humidity,omitempty"`
	DateMaxTemp *int64   `json:"date_max_temp,omitempty"`
	DateMinTemp *int64   `json:"date_min_temp,omitempty"`
	MinTemp     *float64 `json:"min_temp,omitempty"`
	MaxTemp     *float64 `json:"max_temp,omitempty"`
}

type Module struct {
	ID             *string        `json:"_id,omitempty"`
	Type           *string        `json:"type,omitempty"`
	LastMessage    *int64         `json:"last_message,omitempty"`
	LastSeen       *int64         `json:"last_seen,omitempty"`
	DashboardData  *DashboardData `json:"dashboard_data,omitempty"`
	DataType       []string       `json:"data_type,omitempty"`
	ModuleName     *string        `json:"module_name,omitempty"`
	LastSetup      *int64         `json:"last_setup,omitempty"`
	BatteryVP      *int64         `json:"battery_vp,omitempty"`
	BatteryPercent *int64         `json:"battery_percent,omitempty"`
	RfStatus       *int64         `json:"rf_status,omitempty"`
	Firmware       *int64         `json:"firmware,omitempty"`
}

type Device struct {
	ID              *string   `json:"_id,omitempty"`
	CipherID        *string   `json:"cipher_id,omitempty"`
	LastStatusStore *int64    `json:"last_status_store,omitempty"`
	Modules         []*Module `json:"modules,omitempty"`
}

type Body struct {
	Devices []*Device `json:"devices,omitempty"`
}

type StationData struct {
	Body *Body `json:"body,omitempty"`
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
	queries.Add("access_token", d.token)
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
