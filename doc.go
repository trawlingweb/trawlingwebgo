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

// Query function
func Query(params TrwRequest) (TrwResponse, error) {
	v := reflect.ValueOf(params)
	twurl := "https://api.trawlingweb.com/?"
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() != "" {
			if i != 0 {
				twurl += "&"
			}
			if v.Type().Field(i).Name == "Query" {
				sturl := v.Field(i).String()
				encodedPath := url.QueryEscape(sturl)
				twurl += strings.ToLower(v.Type().Field(i).Name) + "=" + encodedPath
			} else {
				twurl += strings.ToLower(v.Type().Field(i).Name) + "=" + v.Field(i).String()
			}
		}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	var res TrwResponse

	req, err := http.NewRequest("GET", twurl, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("User-Agent", "Trawlingweb Go Client 1.0")

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	err2 := json.NewDecoder(resp.Body).Decode(&res)
	return res, err2

}
