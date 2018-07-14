package main

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", getInfo)
	http.ListenAndServe(":8080", nil)
}

// Handler the request and write a response
func getInfo(resp http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()
	apiQuery := buildAPIQuery(queryValues)

	homeInfo, err := callGetSearchResults(apiQuery)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Header().Set(contentTypeKey, xmlContentType)
	xml.NewEncoder(resp).Encode(homeInfo)
}

// Parse the request URL to construct the API query
func buildAPIQuery(queryValues url.Values) string {
	return getSearchResultsURL +
		"?" + zwsIDParam + "=" + zwsID +
		"&" + addressParam + "=" + queryValues.Get(addressParam) +
		"&" + cityStateZipParam + "=" + queryValues.Get(cityStateZipParam)
}

// Calls GetSearchResults API and populates homeInfo with response
func callGetSearchResults(url string) (searchResults, error) {
	resp, apiErr := http.Get(url)
	if apiErr != nil {
		return searchResults{}, apiErr
	}
	defer resp.Body.Close()

	var homeInfo searchResults
	xmlErr := xml.NewDecoder(resp.Body).Decode(&homeInfo)
	if xmlErr != nil {
		return searchResults{}, xmlErr
	}
	return homeInfo, nil
}

const getSearchResultsURL = "http://www.zillow.com/webservice/GetSearchResults.htm"
const zwsIDParam = "zws-id"
const addressParam = "address"
const cityStateZipParam = "citystatezip"
const contentTypeKey = "Content-Type"
const xmlContentType = "application/xml; charset=utf-8"
