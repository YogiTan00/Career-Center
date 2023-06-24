package valueobject

import (
	"strings"
)

type TypeRolesEnum string

const (
	ROLE_MEMBER  TypeRolesEnum = "MEMBER"
	ROLE_ADMIN   TypeRolesEnum = "ADMIN"
	ROLE_COMPANY TypeRolesEnum = "COMPANY"
)

type TypeRoles struct {
	value TypeRolesEnum
}

func NewTypeRoles(s TypeRolesEnum) TypeRoles {
	return TypeRoles{value: s}
}

func (s TypeRoles) StringRoles() string {
	return string(s.value)
}

func NewTypeRolesFromString(s string) TypeRoles {
	var typeRoles TypeRolesEnum
	switch strings.ToUpper(s) {
	case string(ROLE_MEMBER):
		typeRoles = ROLE_MEMBER
	case string(ROLE_ADMIN):
		typeRoles = ROLE_ADMIN
	case string(ROLE_COMPANY):
		typeRoles = ROLE_COMPANY
	default:
		typeRoles = ROLE_MEMBER
	}
	return TypeRoles{value: typeRoles}
}
