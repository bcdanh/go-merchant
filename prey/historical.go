package prey

type HistoricalData struct {
	Date      string
	Open      float32
	High      float32
	Low       float32
	Close     float32
	Volume    int
	MarketCap int
	Average   float32
}

type RealtimeData struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Symbol            string  `json:"symbol"`
	Slug              string  `json:"website_slug"`
	Rank              int     `json:"rank"`
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
	Data              struct {
		Currency struct {
			Price     float32 `json:"price"`
			Volume24h float64 `json:"volume_24h"`
			MarketCap float64 `json:"market_cap"`
			Rate1h    float32 `json:"percent_change_1h"`
			Rate24h   float32 `json:"percent_change_24h"`
			Rate7d    float32 `json:"percent_change_7d"`
		} `json:"USD"`
	} `json:"quotes"`
}

type AllTickerResponse struct {
	Data []RealtimeData `json:"data"`
	Meta Metadata       `json:"metadata"`
}

// type RealtimeData struct {
// 	Time              string
// 	MarketCap         int64
// 	Price             float32
// 	CirculatingSupply int64
// 	Volume24h         int64
// 	Rate1h            float32
// 	Rate24h           float32
// 	Rate7d            float32
// }

// type HistoricalDataList struct {
// 	CoinName string
// 	List     []HistoricalData
// }

// type RealtimeDataList struct {
// 	CoinName []string
// 	List     []RealtimeData
// }
