package prey

type Coin struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	WebsiteSlug string `json:"website_slug"`
}

type Metadata struct {
	Timestamp           int    `json:"timestamp"`
	NumCryptocurrencies int    `json:"num_cryptocurrencies"`
	Error               string `json:"error"`
}

type ListingResponse struct {
	Data []Coin   `json:"data"`
	Meta Metadata `json:"metadata"`
}
