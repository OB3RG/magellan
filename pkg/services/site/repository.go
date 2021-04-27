package site

import (
	"context"
	"database/sql"
	"log"
)

type Site struct {
	ID  int    `json:"id"`
	Url string `json:"url"`
}

func (s *Site) SiteSave() {
	log.Print("Save site")
}

func List(ctx context.Context, db *sql.DB) ([]Site, error) {
	rows, err := db.Query("SELECT * FROM site")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sites []Site
	for rows.Next() {
		site := Site{}
		err = rows.Scan(&site.ID, &site.Url)
		if err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return sites, nil
}

func GetOne(ctx context.Context, db *sql.DB, id string) (*Site, error) {
	row := db.QueryRow("SELECT * FROM site WHERE id=$1", id)
	site := &Site{}

	switch err := row.Scan(&site.ID, &site.Url); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return site, nil
	default:
		return nil, err
	}
}
