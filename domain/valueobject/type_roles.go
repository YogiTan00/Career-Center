package valueobject

import (
	"strings"
)

type TypeRolesEnum string

const (
	MEMBER TypeRolesEnum = "MEMBER"
	ADMIN  TypeRolesEnum = "ADMIN"
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
	case string(MEMBER):
		typeRoles = MEMBER
	case string(ADMIN):
		typeRoles = ADMIN
	default:
		typeRoles = MEMBER
	}
	return TypeRoles{value: typeRoles}
}
