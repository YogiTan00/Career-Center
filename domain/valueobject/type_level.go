package valueobject

import "strings"

type TypeLevelEnum string

const (
	SMA TypeLevelEnum = "SMA"
	SMK TypeLevelEnum = "SMK"
	D3  TypeLevelEnum = "D3"
	S1  TypeLevelEnum = "S1"
	S2  TypeLevelEnum = "S2"
)

type TypeLevel struct {
	value TypeLevelEnum
}

func NewTypeLevel(s TypeLevelEnum) TypeLevel {
	return TypeLevel{value: s}
}

func (s TypeLevel) StringLevel() string {
	return string(s.value)
}

func NewTypeLevelFromString(s string) TypeLevel {
	var typeLevel TypeLevelEnum
	switch strings.ToUpper(s) {
	case string(SMA):
		typeLevel = SMA
	case string(SMK):
		typeLevel = SMK
	case string(D3):
		typeLevel = D3
	case string(S1):
		typeLevel = S1
	case string(S2):
		typeLevel = S2
	}
	return TypeLevel{value: typeLevel}
}
