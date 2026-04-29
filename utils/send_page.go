package utils

import (
	"net/http"
)

type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
}

type PaginatedData struct {
	Content    any        `json:"content"`
	Pagination Pagination `json:"pagination"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, cnt int64) {
	paginatedData := PaginatedData{
		Content: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: cnt,
			TotalPages: cnt / limit,
		},
	}
	SendData(w, true, "Data loaded successfully", paginatedData, http.StatusOK)
}
