package entity

type Pagination struct {
	Page        int `form:"page" json:"page"`
	PageSize    int `form:"page_size" json:"page_size"`
	DataPerPage int `form:"data_per_page" json:"data_per_page"`
	TotalData   int `form:"total_data" json:"total_data"`
}
