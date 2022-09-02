package entities

import (
	"database/sql"
)

// Retrieve from https://openweathermap.org/api.
// Save to DB. Serve w/ endpoints
type WeatherReport struct {
	Temp        int `json:"temp"`
	Humidity    int `json:"humidity,omitempty"`
	Visibility  int `json:"visibility,omitempty"`
	WindSpeed   int `json:"windSpeed,omitempty"`
	WindGust    int `json:"windGust,omitempty"`
	ConditionId int `json:"conditionId,omitempty"`
	UpdatedDate int `json:"updatedOn"`
}

func RetrieveWeatherReportsPerCity(db *sql.DB, cityId string) ([]WeatherReport, error) {
	r, err := db.Query(`
		SELECT temp, humidity, visibility, windSpeed, windGust, updatedDate
		FROM weather_reports
		WHERE cityId = ?`, cityId)
	if err != nil {
		return nil, err
	}

	wrs := make([]WeatherReport, 0, 1000)
	for r.Next() {
		h := sql.NullInt64{}
		v := sql.NullInt64{}
		ws := sql.NullInt64{}
		wg := sql.NullInt64{}
		var d int
		var t int

		err := r.Scan(&t, &h, &v, &ws, &wg, &d)
		if err != nil {
			return nil, err
		}

		wrs = append(wrs, WeatherReport{
			Humidity:    int(h.Int64),
			Visibility:  int(v.Int64),
			Temp:        t,
			WindSpeed:   int(ws.Int64),
			WindGust:    int(wg.Int64),
			UpdatedDate: d,
		})
	}

	return wrs, nil
}
