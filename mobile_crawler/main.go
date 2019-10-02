package main

import (
	"os"

	log "github.com/llimllib/loglevel"
	"github.com/nandotheessen/golang_playground/internal/crawler"
	"github.com/nandotheessen/golang_playground/internal/persistence"
)

func main() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	models := make(map[string]string)
	models["slk-320"] = "https://www.autoscout24.de/lst/mercedes-benz/slk-320/14482-potsdam?sort=standard&desc=0&ustate=N%2CU&size=20&page=1&lon=13.11216&lat=52.39148&zip=14482%20Potsdam&zipr=150&cy=D&atype=C&sort\\=standard\\&desc\\=0\\&ustate\\=N%2CU\\&size\\=20\\&page\\=1\\&lon\\=13.11216\\&lat\\=52.39148\\&zip\\=14482%20Potsdam\\&zipr\\=150\\&cy\\=D\\&fregfrom\\=2000\\&atype\\=C\\&"
	models["slk-32-amg"] = "https://www.autoscout24.de/lst/mercedes-benz/slk-32-amg/14482-potsdam?sort=standard&desc=0&ustate=N%2CU&size=20&page=1&lon=13.11216&lat=52.39148&zip=14482%20Potsdam&zipr=150&cy=D&atype=C&sort\\=standard\\&desc\\=0\\&ustate\\=N%2CU\\&size\\=20\\&page\\=1\\&lon\\=13.11216\\&lat\\=52.39148\\&zip\\=14482%20Potsdam\\&zipr\\=150\\&cy\\=D\\&fregfrom\\=2000\\&atype\\=C\\&"
	models["sl-350"] = "https://www.autoscout24.de/lst/mercedes-benz/sl-350/14482-potsdam?sort=standard&desc=0&ustate=N%2CU&size=20&page=1&lon=13.11216&lat=52.39148&zip=14482%20Potsdam&zipr=150&cy=D&fregto=2008&fregfrom=2006&atype=C&sort\\=standard\\&desc\\=0\\&ustate\\=N%2CU\\&size\\=20\\&page\\=1\\&lon\\=13.11216\\&lat\\=52.39148\\&zip\\=14482%20Potsdam\\&zipr\\=150\\&cy\\=D\\&fregfrom\\=2000\\&atype\\=C\\&"
	models["clk-350"] = "https://www.autoscout24.de/lst/mercedes-benz/clk-350/14482-potsdam?sort=standard&desc=0&ustate=N%2CU&size=20&page=1&lon=13.11216&lat=52.39148&zip=14482%20Potsdam&zipr=150&cy=D&fregto=2010&fregfrom=2006&atype=C&sort\\=standard\\&desc\\=0\\&ustate\\=N%2CU\\&size\\=20\\&page\\=1\\&lon\\=13.11216\\&lat\\=52.39148\\&zip\\=14482%20Potsdam\\&zipr\\=150\\&cy\\=D\\&fregfrom\\=2000\\&atype\\=C\\&"

	// models := []string{"slk-320", "slk-32-amg", "sl-230", "clk-230"}

	Persister := persistence.CreateDatabasePersister(user, password)
	log.SetPriorityString("debug")
	log.SetPrefix("crawler")
	for model, url := range models {
		crawler.FetchPrice(url, model, Persister)
	}
}
