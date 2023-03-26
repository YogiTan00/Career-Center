package valueobject

import "strings"

type TypeCompanyEnum string

const (
	RETAIL         TypeCompanyEnum = "RETAIL"
	MINING         TypeCompanyEnum = "MINING"
	CAPITAL_MARKET TypeCompanyEnum = "CAPITAL_MARKET"
)

type TypeCompany struct {
	value TypeCompanyEnum
}

func NewTypeCompany(s TypeCompanyEnum) *TypeCompany {
	return &TypeCompany{value: s}
}

func (s TypeCompany) StringCompany() string {
	return string(s.value)
}

func NewTypeCompanyFromString(s string) *TypeCompany {
	var typeSearch TypeCompanyEnum
	switch strings.ToUpper(s) {
	case string(RETAIL):
		typeSearch = RETAIL
	case string(MINING):
		typeSearch = MINING
	case string(CAPITAL_MARKET):
		typeSearch = CAPITAL_MARKET
	}
	return &TypeCompany{value: typeSearch}
}
