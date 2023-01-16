package app

type Pager struct {
	Total int64 `json:"total"`
	Data  any   `json:"data"`
}
