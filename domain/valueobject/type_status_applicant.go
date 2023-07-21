package valueobject

import "strings"

type TypeStatusApplicantEnum string

const (
	STATUS_SEND     TypeStatusApplicantEnum = "STATUS_SEND"
	STATUS_PROCCESS TypeStatusApplicantEnum = "STATUS_PROCCESS"
	STATUS_REJECT   TypeStatusApplicantEnum = "STATUS_REJECT"
	STATUS_ACCEPT   TypeStatusApplicantEnum = "STATUS_ACCEPT"
)

type TypeStatusApplicant struct {
	value TypeStatusApplicantEnum
}

//func NewTypeStatusApplicant(s TypeStatusApplicantEnum) TypeStatusApplicant {
//	return TypeStatusApplicant{value: s}
//}

func (s TypeStatusApplicant) String() string {
	return string(s.value)
}

func NewTypeStatusApplicantFromStringPointer(s string) *TypeStatusApplicant {
	var typeStatus TypeStatusApplicantEnum
	switch strings.ToUpper(s) {
	case string(STATUS_SEND):
		typeStatus = STATUS_SEND
	case string(STATUS_PROCCESS):
		typeStatus = STATUS_PROCCESS
	case string(STATUS_REJECT):
		typeStatus = STATUS_REJECT
	case string(STATUS_ACCEPT):
		typeStatus = STATUS_ACCEPT
	}
	return &TypeStatusApplicant{value: typeStatus}
}
