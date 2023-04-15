package rabbitadmin

func CheckPageLimit(page, limit int) (int, int) {
	if page < 1 {
		page = 1
	}
	if limit <= 1 || limit > 200 {
		limit = 50
	}
	return page, limit
}

type PageForm struct {
	Page    int    `json:"page"`  // current page
	Limit   int    `json:"limit"` // page size
	Keyword string `json:"keyword,omitempty"`
}

type PageResult[T any] struct {
	Page       int    `json:"page,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Keyword    string `json:"keyword,omitempty"`
	TotalCount int    `json:"total,omitempty"`
	Items      []T    `json:"items"`
}
