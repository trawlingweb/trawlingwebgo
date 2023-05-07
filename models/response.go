package models

type WorkerResponse struct {
	Response struct {
		Data         []Tweet `json:"data"`
		TotalResults int     `json:"totalResults"`
		RestResults  int     `json:"restResults"`
		Next         string  `json:"next"`
	} `json:"response"`
}

type GetTwitterScrResponse struct {
	Response struct {
		Data         []SCRTweet `json:"data"`
		TotalResults int        `json:"totalResults"`
		RestResults  int        `json:"restResults"`
		Next         string     `json:"next"`
	} `json:"response"`
}

type DeleteWorkerResponse struct {
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

type TrwErrorResponse struct {
	Response struct {
		Error string `json:"error"`
	} `json:"response"`
}
