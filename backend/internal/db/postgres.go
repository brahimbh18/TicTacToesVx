package db

import "database/sql"

func Open(databaseURL string) (*sql.DB, error) {
	_ = databaseURL
	return nil, nil
}
