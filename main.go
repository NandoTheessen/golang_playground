package main

import (

	// "net/http"
	"database/sql"

	"github.com/go-bongo/bongo"
	log "github.com/llimllib/loglevel"
	_ "github.com/mattn/go-sqlite3"
	// "io/ioutil"
	// "regexp"
	// "strconv"
)

// func normalize(s []byte) []byte {
// 	return append(s, make([]byte, 8-len(s))...)
// }

type PricePoint struct {
	bongo.DocumentBase `bson:",inline"`
	Date               string
	Price              float64
}

func main() {
	log.SetPriorityString("debug")
	log.SetPrefix("crawler")

	// if len(os.Args) < 2 {
	// 	log.Fatalln("Missing Url arg")
	// }
	db, err := sql.Open("sqlite3", "./prices.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table pricepoints (id integer not null primary key, date date, price float);
	`
	_, err = db.Exec(sqlStmt)

}
