package request

type RequestApplication struct {
	CompanyId string `json:"companyId"`
}

type RequestAbility struct {
	Ability []string `json:"ability"`
}

type RequestLanguage struct {
	Language []string `json:"language"`
}
