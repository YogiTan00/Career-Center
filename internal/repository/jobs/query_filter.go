package jobs

import (
	"CareerCenter/domain/entity/filter"
	"fmt"
)

func txQuery(f *filter.FilterDTO) string {
	var tx string
	if f != nil {
		tx = fmt.Sprintf("LIMIT %d OFFSET %d", f.Limit, f.Page)
	}
	return tx
}
