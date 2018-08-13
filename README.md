# trawlingweb
Official Javascript Trawlingweb client for Go

[https://trawlingweb.com](https://trawlingweb.com)


### Example
```go
package main

import (
	"fmt"

	trawlingweb "github.com/anpro21/trawlingweb.go"
)

func main() {
	request := trawlingweb.TrwRequest{Token: "ea58ad77426816b16f2cd3c950de07886bc64472", Query: "presidente%20AND%20language:es", Ts: "", Tsi: "", Sort: "", Order: ""}
	ret, err := trawlingweb.Query(request)
	if err != nil {
		fmt.Println((err.Error()))
	} else {
		for _, article := range ret.Response.Data {
			fmt.Println(article.Title)
			fmt.Println(article.URL)
			fmt.Println(article.ID)
			fmt.Println("---------------------")
		}
	}
}

```
