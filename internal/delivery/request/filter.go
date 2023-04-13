package request

import (
	"CareerCenter/domain/entity/filter"
	"net/http"
	"strconv"
)

type RequestFilter struct {
	Q      string `json:"q"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Status bool   `json:"status"`
}

func FilterGeneral(r *http.Request, req *RequestFilter) (*filter.FilterDTO, error) {
	var err error
	q := r.URL.Query().Get("q")
	req.Q = q
	status := r.URL.Query().Get("status")
	if len(status) > 0 {
		req.Status, err = strconv.ParseBool(status)
		if err != nil {
			return nil, err
		}
	}
	limit := r.URL.Query().Get("limit")
	if len(limit) > 0 {
		req.Limit, err = strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
	}
	offset := r.URL.Query().Get("page")
	if len(offset) > 0 {
		req.Page, err = strconv.Atoi(offset)
		if err != nil {
			return nil, err
		}
	}

	return &filter.FilterDTO{
		Q:      req.Q,
		Limit:  uint32(req.Limit),
		Page:   uint32(req.Page),
		Status: req.Status,
	}, nil
}
