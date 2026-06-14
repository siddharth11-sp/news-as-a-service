package news

type ListNewsRequest struct {
	EntityID string

	Sentiment string
	Source    string

	Page     int
	PageSize int
}

type ListNewsResponse struct {
	Data         []NewsArticle `json:"data"`
	TotalRecords int64         `json:"total_records"`
	Page         int           `json:"page"`
	PageSize     int           `json:"page_size"`
}
