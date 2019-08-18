package main

import (

	// "net/http"

	log "github.com/llimllib/loglevel"
	_ "github.com/mattn/go-sqlite3"
	crawler "github.com/nandotheessen/golang_playground/internal/crawler"
	"github.com/nandotheessen/golang_playground/internal/persistence"
	// "io/ioutil"
	// "regexp"
	// "strconv"
)

// func normalize(s []byte) []byte {
// 	return append(s, make([]byte, 8-len(s))...)
// }

type PricePoint struct {
	Date  string
	Price float64
}

func main() {
	log.SetPriorityString("debug")
	log.SetPrefix("crawler")
	persistence.DatabaseConnection()
	// if len(os.Args) < 2 {
	// 	log.Fatalln("Missing Url arg")
	// }

	crawler.FetchPrice()
}
