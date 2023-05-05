package models

type GetWorkerRequest struct {
	Token string
	Ts    string
	Tsi   string
}

type GetArticleRequest struct {
	Token string
	Query string
	Ts    string
	Tsi   string
	Sort  string
	Order string
}
