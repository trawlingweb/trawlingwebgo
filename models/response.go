package models

type WorkerResultsResponse struct {
	Response struct {
		Data         []Tweet `json:"data"`
		TotalResults int     `json:"totalResults"`
		RestResults  int     `json:"restResults"`
		Next         string  `json:"next"`
	} `json:"response"`
}

type TwitterScrResponse struct {
	Response struct {
		Data         []SCRTweet `json:"data"`
		TotalResults int        `json:"totalResults"`
		RestResults  int        `json:"restResults"`
		Next         string     `json:"next"`
	} `json:"response"`
}

type WorkerResponse struct {
	Response struct {
		Worker  string `json:"worker"`
		Message string `json:"msg"`
	} `json:"response"`
}

type ArticleResponse struct {
	Response struct {
		Data         []Article `json:"data"`
		RequestLeft  int       `json:"requestLeft"`
		TotalResults int       `json:"totalResults"`
		RestResults  int       `json:"restResults"`
		Next         string    `json:"next"`
	} `json:"response"`
}

type ErrorResponse struct {
	Response struct {
		Error string `json:"error"`
	} `json:"response"`
}
