package pagination

type SQLPaging struct {
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (paging *SQLPaging) GetOffset() int {
	if paging.Page > 0 && paging.Limit > 0 {
		return (paging.Page - 1) * paging.Limit
	} else {
		return 0
	}
}

var DefaultSQLPaging SQLPaging = SQLPaging{
	Limit: 25,
	Page:  1,
}
