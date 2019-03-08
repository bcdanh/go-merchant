package librarian

import (
	"fmt"
	"log"
	"merchant/prey"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func PushRealtimeData(data prey.RealtimeData) {
	var timestring = string(time.Now().Add(-7*time.Hour).Format("2006-01-02 15:04"))
	//fmt.Println(timestring)
	key := data.ID
	if isNewCoin(data.ID) {
		s := fmt.Sprintf(stmtCreateRealtimeDataTable, key)
		err = taskQuery(s)
		//log.Println(s)
		if err != nil {
			log.Println(s)
			log.Println(err)
		}

		s = fmt.Sprintf(stmtInsertCoinToCoinsTable, data.ID, strings.Replace(data.Name, "'", "", -1), data.Symbol, data.Slug)
		// log.Println(s)
		err = taskQuery(s)
		if err != nil {
			log.Println(s)
			log.Println(err)
		}

	}

	s := fmt.Sprintf(stmtInsertToRealtimeDataTable, key, timestring, data.Rank, data.CirculatingSupply, data.Data.Currency.Price, data.Data.Currency.Volume24h, data.Data.Currency.MarketCap, data.Data.Currency.Rate1h, data.Data.Currency.Rate24h, data.Data.Currency.Rate7d)
	//log.Println(s)
	err = taskQuery(s)
	if err != nil {
		log.Println(s)
		log.Println(err)
	}
}

func isNewCoin(id int) bool {
	for i := 0; i < len(coinsID); i++ {
		if i == id {
			return false
		}
	}
	return true
}
