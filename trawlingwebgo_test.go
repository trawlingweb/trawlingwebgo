package trawlingwebgo

import (
	"testing"

	"github.com/websays-intelligence/trawlingwebgo/models"
)

func TestArticleRequest(t *testing.T) {
	request := models.ArticleRequest{Token: "ea58ad77426816b16f2cd3c950de07886bc64472", Query: "\"Girona\" AND \"Barcelona\""}
	ret, err := Query(request)
	t.Log(ret.Response.TotalResults)
	if err != nil {
		panic(err.Error())
	} else {
		t.Log("Articles: ", len(ret.Response.Data))
		for _, article := range ret.Response.Data {
			t.Log(article.Title)
		}
	}
}

func TestWorkerRequest(t *testing.T) {
	request := models.WorkerRequest{Token: "ea58ad77426816b16f2cd3c950de07886bc64472"}
	ret, err := GetWorker("", request)
	t.Log(ret.Response.TotalResults)
	if err != nil {
		panic(err.Error())
	} else {
		t.Log("Tweets: ", len(ret.Response.Data))
		for _, tweet := range ret.Response.Data {
			t.Log(tweet.PostID)
		}
	}
}
