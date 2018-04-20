package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	githubapi       = "https://api.github.com/search/repositories?q="
	sortpagepartone = "&sort=updated&page="
	pagepartwo      = "&per_page="
)

// Consumehub queries github api for repos and prints. First parameter 'query' is mandatory.
// Next parameters are for pagination with first one for page and next for count of repositories on
// the page
func Consumehub(query string, page int, count int) (data JSONData, err error) {

	if len(query) == 0 {
		err = fmt.Errorf("query cannot be empty")
		return
	}

	q := githubapi + query

	if page > 0 && count > 0 {
		q = q + sortpagepartone + strconv.Itoa(page) + pagepartwo + strconv.Itoa(count)
	}
	res, err := http.Get(q)

	if err != nil {
		err = fmt.Errorf("Consumehub failed, err='%s'", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	data = JSONData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}
