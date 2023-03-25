package jobs

import (
	"CareerCenter/domain/entity/filter"
	"fmt"
	"strings"
)

func txQuery(f *filter.Filter) string {
	var result []string
	if f != nil {
		if len(f.GetQ()) > 0 {
			like := "%" + f.GetQ() + "%"
			tx := fmt.Sprintf("WHERE position like '%s'", like)
			result = append(result, tx)
		}
		if f.GetLimit() > 0 {
			tx := fmt.Sprintf("LIMIT %d", f.GetLimit())
			result = append(result, tx)
		}
		if f.GetPage() > 0 {
			tx := fmt.Sprintf("OFFSET %d", f.GetPage())
			result = append(result, tx)
		}
	}
	tx := strings.Join(result, " ")
	return tx
}
