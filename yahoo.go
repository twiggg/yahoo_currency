package yahoo

import (
	"fmt"
	//"net/url"
	"io/ioutil"
	"net/http"
	"strings"
)

var currencies_map = map[string]string{
	//asia pacific
	"australia":   "AUD",
	"china":       "CNY",
	"hong kong":   "HKD",
	"indonesia":   "IDR",
	"japan":       "JPY",
	"korea":       "KRW",
	"malaysia":    "MYR",
	"new zealand": "NZD",
	"philippines": "PHP",
	"singapore":   "SGD",
	"taiwan":      "TWD",
	"thailand":    "THB",
	"vietnam":     "VND",
	//europe middle east
	"belgium":       "EUR",
	"czek republic": "CZK",
	"danmark":       "EUR",
	"germany":       "EUR",
	"spain":         "EUR",
	"france":        "EUR",
	"hungary":       "HUF",
	"ireland":       "EUR",
	"italia":        "EUR",
	"luxembourg":    "EUR",
	"netherland":    "EUR",
	"norway":        "NOK",
	"austria":       "EUR",
	"poland":        "PLN",
	"portugal":      "EUR",
	"russia":        "RUB",
	"swiss":         "CHF",
	"sweden":        "SEK",
	"finland":       "EUR",
	"turkey":        "TRY",
	"uk":            "GBP",
	"uae":           "AED",
	//americas
	"brazil": "BRL",
	"canada": "CAD",
	"mexico": "MXN",
	"usa":    "USD",
}

func uniqueFromMapStringString(m map[string]string) map[string]int {
	currencies_set := map[string]int{}

	for _, curr := range m {
		if _, ok := currencies_set[curr]; !ok {
			currencies_set[curr] = 1
		} else {
			currencies_set[curr] = currencies_set[curr] + 1
		}
	}
	//fmt.Println("list of currencies")
	/*for curr, nb := range currencies_set {
		fmt.Println(curr, " , ", nb)
	}*/
	return currencies_set
}

func currenciesCombos(m map[string]int, ref string) string {
	combos := []string{}
	for curr, _ := range m {
		if curr != ref {
			combos = append(combos, fmt.Sprintf("'%s%s'", ref, curr))
		}
	}
	//fmt.Println("number of currencies : ",len(m))
	//fmt.Println("number of combos : ",len(combos))
	query := ""
	for i, v := range combos {

		if i == 0 {
			query = v
		} else {
			query = query + "," + v
		}

	}
	//fmt.Println("length of query : ",len(query))
	if len(query) != len(combos)*(6+2)+len(combos)-1 {
		fmt.Println("the query is incorrect")
	}
	return query
}

func formulateYQL(pairs string) string {
	part1 := "select * from yahoo.finance.xchange where pair in ("
	part2 := pairs
	part3 := ")"

	query := part1 + part2 + part3
	return query

}

func UpdateQuery() string {
	fmt.Println("updating query string for yahoo api...")
	unique := uniqueFromMapStringString(currencies_map)
	q := currenciesCombos(unique, "EUR")
	//fmt.Println(q)
	q2 := formulateYQL(q)
	//fmt.Println("yahoo query :")
	//fmt.Println(q2)
	q3 := strings.Replace(q2, " ", "%20", -1)
	query := strings.Replace(q3, ",", "%2C", -1)
	fmt.Println("escaped query")
	fmt.Println(query)
	return query

}

func GetFromYahoo() ([]byte, error) {
	query := "select%20*%20from%20yahoo.finance.xchange%20where%20pair%20in%20('EURGBP'%2C'EURCHF'%2C'EURHUF'%2C'EURMYR'%2C'EURUSD'%2C'EURCNY'%2C'EURNZD'%2C'EURSEK'%2C'EURTRY'%2C'EURMXN'%2C'EURJPY'%2C'EURPLN'%2C'EURAUD'%2C'EURCZK'%2C'EURAED'%2C'EURSGD'%2C'EURTWD'%2C'EURHKD'%2C'EURPHP'%2C'EURBRL'%2C'EURCAD'%2C'EURVND'%2C'EURNOK'%2C'EURIDR'%2C'EURRUB'%2C'EURTHB'%2C'EURKRW')"
	starturl := "https://query.yahooapis.com/v1/public/yql?q="
	endurl := "&format=json&diagnostics=true&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys&callback="
	url := starturl + query + endurl
	//make the call for yahoo api
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("there was an error while making yahoo api call : ", err)
		return []byte(""), err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("there was an error while reading response's body : ", err)
			return []byte(""), err
		} else {
			//fmt.Printf(string(body))
			return body, nil
		}
	}

}

/*
https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20yahoo.finance.xchange%20where%20pair%20in%20('EURTHB'%2C'EURAUD'%2C'EURUSD'%2C'EURIDR'%2C'EURCZK'%2C'EURMXN'%2C'EURVND'%2C'EURNZD'%2C'EURTWD'%2C'EURHUF'%2C'EURSEK'%2C'EURHKD'%2C'EURCAD'%2C'EURMYR'%2C'EURTRY'%2C'EURJPY'%2C'EURPLN'%2C'EURKRW'%2C'EURNOK'%2C'EURPHP'%2C'EURRUB'%2C'EURCHF'%2C'EURCNY'%2C'EURGBP'%2C'EURAED'%2C'EURBRL'%2C'EURSGD')&format=json&diagnostics=true&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys&callback=
*/
