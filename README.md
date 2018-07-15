# :house: ZHomeInfo
## Tasks

- [x] **Learn** [Go](https://golang.org/) syntax, principles
- [x] Write **backend** Go app that spins up an http server on 8080
- [x] Parse the request URL for **address**, and **city/state or zipcode**
- [x] **Call** [Zillow's API](https://www.zillow.com/howto/api/GetSearchResults.htm) with the query parameters
- [x] **Decode** the XML response into a structure
- [ ] Write **frontend** code to accept input
- [ ] Write **frontend** code to display output
- [ ] **Host** web app publicly on Heroku or similar
- [x] Put **source code** up on Github
- [x] Create **README**

## Known Issues

- `xml.NewDecoder().Decode()` doesn't throw decoding errors if the API response structure is not as expected

## Notes

- This project was put on hold after one night's work.

## How to Run

1. [Install Go](https://golang.org/doc/install)

2. Download the source files for this project and save them to `<your Go workspace>/src/ZHomeInfo`.

3. `cd ZHomeInfo`

4. Create a file called `config.go` and populate it with your API key:

   1. ```go
      package main
      
      const zwsID = "<your zwsID>"
      ```

5. `go build`

6. `./ZHomeInfo`

7. In another terminal, `curl "http://localhost:8080/?address=<your address>&citystatezip=<your city and state or zip>"`