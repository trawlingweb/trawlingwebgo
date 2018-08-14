package trawlingwebgo

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// TrwResponse API structure
type TrwResponse struct {
	Response struct {
		Data []struct {
			SiteRegion   string    `json:"site_region"`
			SiteType     string    `json:"site_type"`
			Crawled      int64     `json:"crawled"`
			Author       string    `json:"author"`
			Language     string    `json:"language"`
			Published    time.Time `json:"published"`
			Title        string    `json:"title"`
			URL          string    `json:"url"`
			Site         string    `json:"site"`
			SiteCountry  string    `json:"site_country"`
			Domain       string    `json:"domain"`
			Text         string    `json:"text"`
			SiteLanguage string    `json:"site_language"`
			ID           string    `json:"id"`
		} `json:"data"`
		RequestLeft  int    `json:"requestLeft"`
		TotalResults int    `json:"totalResults"`
		RestResults  int    `json:"restResults"`
		Next         string `json:"next"`
	} `json:"response"`
}

// TrwRequest API structure
type TrwRequest struct {
	Token string
	Query string
	Ts    string
	Tsi   string
	Sort  string
	Order string
}

// TrwError for get problems
type TrwError struct {
	Response struct {
		Error string `json:"error"`
	} `json:"response"`
}

// Request to https service
func Request(url string) (TrwResponse, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	var res TrwResponse
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("User-Agent", "trawlingweb-cli.go 1.0")
	resp, err2 := client.Do(req)
	if err2 != nil {
		return res, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	err3 := json.NewDecoder(resp.Body).Decode(&res)
	return res, err3
}

// Query Initial function
func Query(params TrwRequest) (TrwResponse, error) {
	values := reflect.ValueOf(params)
	twurl := "https://api.trawlingweb.com/?"
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() != "" {
			if i != 0 {
				twurl += "&"
			}
			if values.Type().Field(i).Name == "Query" {
				sturl := values.Field(i).String()
				encodedPath := url.PathEscape(sturl)
				twurl += "q=" + encodedPath
			} else {
				twurl += strings.ToLower(values.Type().Field(i).Name) + "=" + values.Field(i).String()
			}
		}
	}

	return Request(twurl)
}

// Next query function
func Next(twurl string) (TrwResponse, error) {
	return Request(twurl)
}
