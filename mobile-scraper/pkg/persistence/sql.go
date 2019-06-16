package persistence

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func connect() (db *sql.DB, err error){
	db, err = sql.Open("postgres", "user=theUser dbname=theDbName sslmode=verify-full")
	defer db.Close()
	if err != nil {
		return nil, err
	}
	return db, nil
}