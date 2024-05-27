package entity

type Response struct {
	Message        string      `json:"message"`
	PaginationData *Pagination `json:"pagination,omitempty"`
	Data           any         `json:"data,omitempty"`
	Error          string      `json:"error,omitempty"`
}
