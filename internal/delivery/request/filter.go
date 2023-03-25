package request

import (
	"CareerCenter/domain/entity/filter"
	"net/http"
	"strconv"
)

type RequestFilter struct {
	Q     string `json:"q"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

func FilterGeneral(r *http.Request, req *RequestFilter) (*filter.FilterDTO, error) {
	var err error
	q := r.URL.Query().Get("q")
	req.Q = q
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
		Q:     q,
		Limit: uint32(req.Limit),
		Page:  uint32(req.Page),
	}, nil
}
