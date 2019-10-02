package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/llimllib/loglevel"
)

type Persister struct {
	db *sql.DB
}

func CreateDatabasePersister(user string, password string) *Persister {
	Persister := &Persister{}
	connectStr := fmt.Sprintf("user=postgres password=%v sslmode=disable host=pricesdb.ci0oftugfryf.us-east-1.rds.amazonaws.com", password)
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		log.Fatal("Failure establishing database conection \n", err)
	}

	createDB := `CREATE TABLE pricepoints ( id SERIAL PRIMARY KEY, date INTEGER, model VARCHAR NOT NULL, price FLOAT);`
	_, err = db.Exec(createDB)
	if err != nil && err.Error() != "pq: relation \"pricepoints\" already exists" {
		log.Fatal("db not created\n", err)
	}

	Persister.db = db
	return Persister
}

func (p *Persister) SaveToDB(date int64, model string, price float64) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Fatal("error in db Begin\n", err)
	}
	stmtString := fmt.Sprintf("INSERT INTO pricepoints VALUES(DEFAULT, %v, '%s', %f)", date, model, price)
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
