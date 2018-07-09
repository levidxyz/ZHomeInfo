package main

import (
	"encoding/xml"
	"net/http"
)

func main() {
	http.HandleFunc("/", getInfo)     // install handler func at root of the webserver
	http.ListenAndServe(":8080", nil) // start HTTP server on port 8080
}

// Handler func to handle the request and write a response
func getInfo(resp http.ResponseWriter, req *http.Request) {

	apiQuery := getAPIQueryFromURL(req)
	info, err := callGetSearchResults(apiQuery)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Header().Set("Content-Type", "application/xml; charset=utf-8")
	xml.NewEncoder(resp).Encode(info)
}

func getAPIQueryFromURL(req *http.Request) string {

	queryValues := req.URL.Query()
	address := queryValues.Get("address")
	cityStateZip := queryValues.Get("citystatezip")
	zwsID := "redacted!!" // TODO get safely from config

	return "http://www.zillow.com/webservice/GetSearchResults.htm" +
		"?zws-id=" + zwsID +
		"&address=" + address +
		"&citystatezip=" + cityStateZip
}

// Calls GetSearchResults API and populates homeInfo with response
func callGetSearchResults(url string) (homeInfo, error) {

	resp, apiErr := http.Get(url)
	if apiErr != nil {
		return homeInfo{}, apiErr
	}
	defer resp.Body.Close() // Close the response body whether or not we can decode the XML

	var info homeInfo
	xmlErr := xml.NewDecoder(resp.Body).Decode(&info)
	if xmlErr != nil { // TODO either this doesn't catch decoding errors, or Decode() way too lenient
		return homeInfo{}, xmlErr
	}
	return info, nil
}

// Struct to represent GetSearchResults API response
type homeInfo struct {
	Message struct {
		Text string `xml:"text"`
	} `xml:"message"`

	Response struct {
		Results []struct {
			Zpid string `xml:"zpid"`

			Links struct {
				HomeDetails   string `xml:"homedetails"`
				GraphsAndData string `xml:"graphsanddata"`
				MapThisHome   string `xml:"mapthishome"`
				Comparables   string `xml:"comparables"`
			} `xml:"links"`

			Address struct {
				Street    string `xml:"street"`
				ZipCode   string `xml:"zipcode"`
				City      string `xml:"city"`
				State     string `xml:"state"`
				Latitude  string `xml:"latitude"`
				Longitude string `xml:"longitude"`
			} `xml:"address"`

			Zestimate struct {
				Amount struct {
					Currency string `xml:"currency,attr"`
					Value    string `xml:",chardata"`
				} `xml:"amount"`
				LastUpdated string `xml:"last-updated"`
				ValueChange struct {
					Duration int    `xml:"duration,attr"`
					Currency string `xml:"currency,attr"`
					Value    string `xml:",chardata"`
				} `xml:"valueChange"`
				ValuationRange struct {
					Low struct {
						Currency string `xml:"currency,attr"`
						Value    string `xml:",chardata"`
					} `xml:"low"`
					High struct {
						Currency string `xml:"currency,attr"`
						Value    string `xml:",chardata"`
					} `xml:"high"`
				} `xml:"valuationRange"`
				Percentile int `xml:"percentile"`
			} `xml:"zestimate"`

			LocalRealEstate struct {
				Region []struct {
					ZindexValue         string  `xml:"zindexValue"`
					ZindexOneYearChange float32 `xml:"zindexOneYearChange"`
					Links               struct {
						Overview       string `xml:"overview"`
						ForSaleByOwner string `xml:"forSaleByOwner"`
						ForSale        string `xml:"forSale"`
					} `xml:"links"`
					ID   int    `xml:"id,attr"`
					Type string `xml:"type,attr"`
					Name string `xml:"name,attr"`
				} `xml:"region"`
			} `xml:"localRealEstate"`
		} `xml:"results>result"`
	} `xml:"response"`
}
