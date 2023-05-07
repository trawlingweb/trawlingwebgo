package trawlingwebgo

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/trawlingweb/trawlingwebgo/models"
)

// WorkerRequest to https service
func WorkerRequest(url string) (models.GetWorkerResponse, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	var res models.GetWorkerResponse
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("User-Agent", "trawlingweb-cli.go 1.2")
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

// GetWorker
func GetWorker(workerID string, params models.GetWorkerRequest) (models.GetWorkerResponse, error) {
	values := reflect.ValueOf(params)
	trwurl := fmt.Sprintf("https://api.trawlingweb.com/posts/%v?", workerID)
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() != "" {
			if i != 0 {
				trwurl += "&"
			}
			trwurl += strings.ToLower(values.Type().Field(i).Name) + "=" + values.Field(i).String()
		}
	}

	return WorkerRequest(trwurl)
}

// DeleteWorker
func DeleteWorker(workerID, token string) (models.DeleteWorkerResponse, error) {
	url := fmt.Sprintf("https://api.trawlingweb.com/delete/%v?token=%v", workerID, token)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	var res models.DeleteWorkerResponse
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("User-Agent", "trawlingweb-cli.go 1.2")
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

// WorkerNext query function
func WorkerNext(trwurl string) (models.GetWorkerResponse, error) {
	return WorkerRequest(trwurl)
}

// ArticleRequest to https service
func ArticleRequest(url string) (models.GetArticleResponse, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	var res models.GetArticleResponse
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("User-Agent", "trawlingweb-cli.go 1.2")
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

// ArticleRequest to https service
func SCRTweetRequest(url string) (models.GetTwitterScrResponse, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	var res models.GetTwitterScrResponse
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("User-Agent", "trawlingweb-cli.go 1.2")
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

func TwitterQuery(params models.GetArticleRequest) (models.GetTwitterScrResponse, error) {
	values := reflect.ValueOf(params)
	twurl := "https://twitter.scr.trawlingweb.com/posts/?"
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() != "" {
			if i != 0 {
				twurl += "&"
			}
			if values.Type().Field(i).Name == "Query" {
				sturl := values.Field(i).String()
				encodedPath := url.QueryEscape(sturl)
				twurl += "q=" + encodedPath
			} else {
				twurl += strings.ToLower(values.Type().Field(i).Name) + "=" + values.Field(i).String()
			}
		}
	}

	return SCRTweetRequest(twurl)
}


// Query Initial function
func Query(params models.GetArticleRequest) (models.GetArticleResponse, error) {
	values := reflect.ValueOf(params)
	twurl := "https://api.trawlingweb.com/?"
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() != "" {
			if i != 0 {
				twurl += "&"
			}
			if values.Type().Field(i).Name == "Query" {
				sturl := values.Field(i).String()
				encodedPath := url.QueryEscape(sturl)
				twurl += "q=" + encodedPath
			} else {
				twurl += strings.ToLower(values.Type().Field(i).Name) + "=" + values.Field(i).String()
			}
		}
	}

	return ArticleRequest(twurl)
}

// Next query function
func Next(twurl string) (models.GetArticleResponse, error) {
	return ArticleRequest(twurl)
}

// Next query function
func NextSCRTweet(twurl string) (models.GetTwitterScrResponse, error) {
	return SCRTweetRequest(twurl)
}
