package models

type WorkerRequest struct {
	Token string
	Ts    string
	Tsi   string
}

type ArticleRequest struct {
	Token string
	Query string
	Ts    string
	Tsi   string
	Sort  string
	Order string
}
