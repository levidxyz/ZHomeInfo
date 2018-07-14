package main

// Struct to represent GetSearchResults API response
type searchResults struct {
	Message struct {
		Text string `xml:"text"`
	} `xml:"message"`
	Response response `xml:"response"`
}

type response struct {
	Results []struct {
		Zpid            string          `xml:"zpid"`
		Links           links           `xml:"links"`
		Address         address         `xml:"address"`
		Zestimate       zestimate       `xml:"zestimate"`
		LocalRealEstate localRealEstate `xml:"localRealEstate"`
	} `xml:"results>result"`
}

type links struct {
	HomeDetails   string `xml:"homedetails"`
	GraphsAndData string `xml:"graphsanddata"`
	MapThisHome   string `xml:"mapthishome"`
	Comparables   string `xml:"comparables"`
}

type address struct {
	Street    string `xml:"street"`
	ZipCode   string `xml:"zipcode"`
	City      string `xml:"city"`
	State     string `xml:"state"`
	Latitude  string `xml:"latitude"`
	Longitude string `xml:"longitude"`
}

type zestimate struct {
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
}

type localRealEstate struct {
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
}
