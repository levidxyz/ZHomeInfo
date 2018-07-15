package main

import (
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", getInfo)
	http.ListenAndServe(":8080", nil)
}

// Handler the request and write a response
func getInfo(response http.ResponseWriter, request *http.Request) {
	apiRequest, queryError := buildAPIRequest(request.URL.Query())
	if queryError != nil {
		http.Error(response, queryError.Error(), http.StatusUnprocessableEntity)
		return
	}

	apiResponse, apiError := http.Get(apiRequest)
	if apiError != nil {
		http.Error(response, apiError.Error(), http.StatusInternalServerError)
		return
	}

	homeInfo, xmlError := decodeAPIResponse(apiResponse)
	if xmlError != nil {
		http.Error(response, xmlError.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set(contentTypeKey, xmlContentType)
	xml.NewEncoder(response).Encode(homeInfo)
}

// Parse the request URL to construct the API query
func buildAPIRequest(queryValues url.Values) (string, error) {
	address := queryValues.Get(addressParam)
	cityStateZip := queryValues.Get(cityStateZipParam)
	if address == "" || cityStateZip == "" {
		return "", errors.New("required query parameters are missing")
	}

	return getSearchResultsURL +
			"?" + zwsIDParam + "=" + zwsID +
			"&" + addressParam + "=" + address +
			"&" + cityStateZipParam + "=" + cityStateZip,
		nil
}

// Decode the GetSearchResults API  XML response
func decodeAPIResponse(resp *http.Response) (searchResults, error) {
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
