package hunter

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
)

var regxrank = regexp.MustCompile(`<a class="currency-name-container link-secondary" href="\/vi\/currencies\/(.+)\/">`)
var regxbegin = regexp.MustCompile(`<div class="table-responsive">`)
var regxend = regexp.MustCompile(`<\/tr>\n<\/tbody>\n<\/table>\n<\/div>`)

var rankURL = "https://coinmarketcap.com/vi/"
var historicalURL = "https://coinmarketcap.com/currencies/%s/historical-data/?start=%s&end=%s"

func getRank(n int) []string {
	resprank, err := http.Get(rankURL)
	if err != nil {
		fmt.Println(err)
	}
	defer resprank.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resprank.Body)

	list := regxrank.FindAllSubmatch(buf.Bytes(), -1)
	listcoin := []string{}

	for i := 0; i < n && i < 100; i++ {
		listcoin = append(listcoin, string(list[i][1]))
	}
	return listcoin
}

func getHistorical(name string, start string, end string) []byte {
	data := []byte{}
	fmt.Println(fmt.Sprintf(historicalURL))
	resp, err := http.Get(fmt.Sprintf(historicalURL, name, start, end))
	if err != nil {
		fmt.Println(err)
		return data
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	body := buf.Bytes()

	tableFirstIndex := regxbegin.FindIndex(body)
	tableLastIndex := regxend.FindIndex(body)

	data = body[tableFirstIndex[0]:tableLastIndex[len(tableLastIndex)-1]]

	return data
}
