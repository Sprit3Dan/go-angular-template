package entities

import (
	"database/sql"
	"fmt"
)

type City struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
	Id        string  `json:"id"`
}

func RetrieveCities(db *sql.DB) ([]City, error) {
	rows, err := db.Query("SELECT Id, Name FROM cities;")
	if err != nil {
		return nil, err
	}

	cities := make([]City, 0, 1000)
	for rows.Next() {
		c := City{}
		err := rows.Scan(&c.Id, &c.Name)
		if err != nil {
			return nil, err
		}

		cities = append(cities, c)
	}

	return cities, nil
}

func RetrieveCity(db *sql.DB, id string) (*City, error) {
	rows, err := db.Query("SELECT Id, Name FROM cities WHERE Id = ?", id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("City %v not found", id)
	}

	c := City{}
	rows.Scan(&c.Id, &c.Name)

	return &c, nil
}

func CreateCity(db *sql.DB, c *City) (*City, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	tx.Exec("INSERT INTO cities (Id, Name) VALUES (?, ?)", c.Id, c.Name)
	txErr := tx.Commit()
	if txErr != nil {
		return nil, txErr
	}

	createdCity, err := RetrieveCity(db, c.Id)

	return createdCity, err
}
