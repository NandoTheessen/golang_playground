package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"

	log "github.com/llimllib/loglevel"
)

func FetchPrice() {
	var prices []float64
	url := os.Args[1]
	fmt.Println("Fetching HTML...")
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Fatalln("Error fetching URL!")
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading response body!")
	}
	log.Infoln("Creating regex's")
	priceMatcher, _ := regexp.Compile("Preis: € (\\d+.\\d+)")
	priceExtractor, _ := regexp.Compile("(\\d+\\.\\d+)")
	result := priceMatcher.FindAll(html, -1)
	if result == nil {
		log.Infoln("No cars available!")
	}
	for _, priceString := range result {
		priceInBytes := priceExtractor.Find(priceString)
		output, err := strconv.ParseFloat(fmt.Sprintf("%s", priceInBytes), 10)
		if err != nil {
			log.Fatalln("Price format changed!")
		}
		log.Infoln(output * 1000)
		prices = append(prices, output*1000)
	}

}
