package main

import (
	"os"

	log "github.com/llimllib/loglevel"
	"github.com/nandotheessen/golang_playground/internal/crawler"
	"github.com/nandotheessen/golang_playground/internal/persistence"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: crawler <model> <queryURL>")
	}

	url := os.Args[2]
	model := os.Args[1]
	Persister := persistence.CreateDatabasePersister()
	log.SetPriorityString("debug")
	log.SetPrefix("crawler")

	crawler.FetchPrice(url, model, Persister)
}
