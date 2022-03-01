# trawlingweb
Official Trawlingweb client for Go Language

[https://trawlingweb.com](https://trawlingweb.com)


### Example
```go
package main

import (
	"fmt"

	trw "github.com/anpro21/trawlingwebgo"
)

func main() {
	request := trw.TrwRequest{Token: "ea58ad77426816b16f2cd3c950de07886bc64472", Query: "debian AND redhat", Ts: "1533459254944"}
	ret, err := trw.Query(request)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Articles: ", len(ret.Response.Data))
		for _, article := range ret.Response.Data {
			fmt.Println(article.Title)
			fmt.Println(article.URL)
			fmt.Println(article.ID)
			fmt.Println("---------------------")
		}
	}

	urlnext := ret.Response.Next
	for urlnext != "" {
		ret, err = trw.Next(urlnext)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Articles: ", len(ret.Response.Data))
		fmt.Println("Rest results: ", ret.Response.RestResults)
		fmt.Println("---------------------")
		urlnext = ret.Response.Next
	}

	fmt.Println("End Query")

}
```

### Links
[Register](https://dashboard.trawlingweb.com/register)

[Control Panel](https://dashboard.trawlingweb.com)

[API documentation](https://dashboard.trawlingweb.com/dashboard)

[Github Repository](https://github.com/anpro21)

### Functions
```
Query(params TrwRequest)
Next(nexturl string)
```

### Input Params
* Token: string with service key
* Query: String with trawling query
* Ts: unixtimestamp in ms
* Tsi: unixtimestamp in ms
* Sort: "published"/"crawled"
* Order: "asc"/"desc"

### Input Params struct

```go
type TrwRequest struct {
	Token string
	Query string
	Ts    string
	Tsi   string
	Sort  string
	Order string
}
```

### Return struct
```go
type TrwResponse struct {
	Response struct {
		Data         []TrwArticle `json:"data"`
		RequestLeft  int          `json:"requestLeft"`
		TotalResults int          `json:"totalResults"`
		RestResults  int          `json:"restResults"`
		Next         string       `json:"next"`
	} `json:"response"`
}
```

### Article struct
```go
type TrwArticle struct {
	Crawled      int64     `json:"crawled"`
	Author       string    `json:"author"`
	Language     string    `json:"language"`
	Published    time.Time `json:"published"`
	Title        string    `json:"title"`
	URL          string    `json:"url"`
	Section      string    `json:"section"`
	Value        float64   `json:"value"`
	Rank         int64     `json:"rank"`
	Visitors     int64     `json:"unique_visitors"`
	Site         string    `json:"site"`
	SiteType     string    `json:"site_type"`
	SiteSection  string    `json:"site_section"`
	SiteLanguage string    `json:"site_language"`
	SiteRegion   string    `json:"site_region"`
	SiteCountry  string    `json:"site_country"`
	Domain       string    `json:"domain"`
	Text         string    `json:"text"`
	ID           string    `json:"id"`
}
```


### MIT License

Copyright (c) 2018 Anpro21

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
