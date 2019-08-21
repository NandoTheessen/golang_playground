package persistence

import (
	"database/sql"
	"fmt"

	log "github.com/llimllib/loglevel"
	_ "github.com/mattn/go-sqlite3"
)

type Persister struct {
	db *sql.DB
}

func CreateDatabasePersister() *Persister {
	Persister := &Persister{}
	db, err := sql.Open("sqlite3", "./prices.db")
	if err != nil {
		log.Fatal("Failure establishing database conection \n", err)
	}

	sqlStmt := `
	create table pricepoints (id integer primary key autoincrement, date integer, model text, price int);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil && err.Error() != "table pricepoints already exists" {
		log.Fatal("db not created", err)
	}

	Persister.db = db
	return Persister
}

func (p *Persister) SaveToDB(date int64, model string, price float64) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Fatal("error in db Begin\n", err)
	}
	stmtString := fmt.Sprintf("INSERT INTO pricepoints VALUES(null, %v, '%s', %f)", date, model, price)
	fmt.Println(stmtString)
	stmt, err := tx.Prepare(stmtString)
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
