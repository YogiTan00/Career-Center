package request

import (
	"CareerCenter/domain/entity/filter"
	"net/http"
	"strconv"
)

type RequestFilter struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func FilterGeneral(r *http.Request, req *RequestFilter) (*filter.FilterDTO, error) {
	var err error
	limit := r.URL.Query().Get("limit")
	req.Limit, err = strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}
	offset := r.URL.Query().Get("page")
	req.Page, err = strconv.Atoi(offset)
	if err != nil {
		return nil, err
	}
	return &filter.FilterDTO{
		Limit: uint32(req.Limit),
		Page:  uint32(req.Page),
	}, nil
}
