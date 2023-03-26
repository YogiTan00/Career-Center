package valueobject

import "strings"

type TypeSearchEnum string

const (
	JOBS    TypeSearchEnum = "JOBS"
	COMPANY TypeSearchEnum = "COMPANY"
)

type TypeSearch struct {
	value TypeSearchEnum
}

func NewTypeSearch(s TypeSearchEnum) *TypeSearch {
	return &TypeSearch{value: s}
}

func (s TypeSearch) StringSearch() string {
	return string(s.value)
}

func NewTypeSearchFromString(s string) *TypeSearch {
	var typeSearch TypeSearchEnum
	switch strings.ToUpper(s) {
	case string(JOBS):
		typeSearch = JOBS
	case string(COMPANY):
		typeSearch = COMPANY
	}
	return &TypeSearch{value: typeSearch}
}
