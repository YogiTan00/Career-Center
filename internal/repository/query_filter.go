package repository

import (
	"CareerCenter/domain/entity/filter"
	"CareerCenter/domain/valueobject"
	"fmt"
	"strings"
)

func TxQuery(typeSearch *valueobject.TypeSearch, f *filter.Filter) string {
	var result []string
	if f != nil {
		if len(f.GetQ()) > 0 {
			like := "%" + f.GetQ() + "%"
			if typeSearch.StringSearch() == string(valueobject.JOBS) {
				tx := fmt.Sprintf("WHERE position like '%s'", like)
				result = append(result, tx)
			} else if typeSearch.StringSearch() == string(valueobject.COMPANY) {
				tx := fmt.Sprintf("WHERE name like '%s'", like)
				result = append(result, tx)
			} else {
				fmt.Println("error type search")
			}
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
