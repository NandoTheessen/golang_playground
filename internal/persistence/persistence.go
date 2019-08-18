package persistence

import (
	"database/sql"

	log "github.com/llimllib/loglevel"
	_ "github.com/mattn/go-sqlite3"
)

func DatabaseConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./prices.db")
	if err != nil {
		log.Fatal("Failure establishing database conection \n", err)
	}
	defer db.Close()

	sqlStmt := `
	create table pricepoints (id integer primary key autoincrement, date date, price float);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil && err.Error() != "table pricepoints already exists" {
		log.Fatal("db not created", err)
	}
	return db
}

func SaveToDB() {
	db := DatabaseConnection()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("error in db Begin\n", err)
	}
	stmt, err := tx.Prepare("INSERT INTO pricepoints VALUES(null, '10.10.1987', '24.90')")
	if err != nil {
		log.Fatal("error in db prepare\n", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
