package hunter

import (
	"encoding/json"
	"fmt"
	"log"
	"merchant/prey"
	"net/http"
)

func GetCoinList() []prey.Coin {
	const requestURL = "https://api.coinmarketcap.com/v2/listings/"
	var listing prey.ListingResponse

	resp, err := http.Get(requestURL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	jsonParser := json.NewDecoder(resp.Body)
	jsonParser.Decode(&listing)

	return listing.Data
}

func GetAllTicker() []prey.RealtimeData {
	const requestURL = "https://api.coinmarketcap.com/v2/ticker/?start=%d&limit=100&sort=id&structure=array"
	var listTicker prey.AllTickerResponse
	for i := 1; i < 1700; i = i + 100 {
		resp, err := http.Get(fmt.Sprintf(requestURL, i))
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		var list prey.AllTickerResponse
		jsonParser := json.NewDecoder(resp.Body)
		jsonParser.Decode(&list)

		listTicker.Data = append(listTicker.Data, list.Data...)
	}

	return listTicker.Data
}

func GetTickers(start int, limit int) []prey.RealtimeData {
	const requestURL = "https://api.coinmarketcap.com/v2/ticker/?start=%d&limit=%d&sort=id&structure=array"
	resp, err := http.Get(fmt.Sprintf(requestURL, start, limit))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	var list prey.AllTickerResponse
	jsonParser := json.NewDecoder(resp.Body)
	jsonParser.Decode(&list)

	return list.Data
}
