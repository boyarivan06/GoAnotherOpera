package api

type Response[T any] struct {
	Headers struct {
		Status       string `json:"status"`
		Code         int    `json:"code"`
		ErrorMessage string `json:"error_message"`
		ResultsCount int    `json:"results_count"`
	} `json:"headers"`
	Results []T `json:"results"`
}
