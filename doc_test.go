package trawlingwebgo

import (
	"testing"
)

func TestRequest(t *testing.T) {
	request := TrwRequest{Token: "ea58ad77426816b16f2cd3c950de07886bc64472", Query: "\"Girona\" AND \"Barcelona\""}
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
