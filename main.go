package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Cred struct {
	Apikey string `json:"apikey"`
}

func main() {

	url := "https://api.yelp.com/v3/businesses/search"

	// get cred from json
	f, err := os.ReadFile("cred.json")
	if err != nil {
		log.Fatal(err)
	}
	var cred Cred
	err = json.Unmarshal(f, &cred)
	if err != nil {
		log.Fatal(err)
	}
	apikey := cred.Apikey

	// Create http request
	req, _ := http.NewRequest("GET", url, nil)

	// Add header
	bearer := "Bearer " + apikey
	req.Header.Add("Authorization", bearer)

	// Add params
	params := req.URL.Query()
	params.Add("latitude", "35.6905")
	params.Add("longitude", "139.6995")
	params.Add("locale", "ja_JP")
	req.URL.RawQuery = params.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp)
	fmt.Println(string(body))

}
