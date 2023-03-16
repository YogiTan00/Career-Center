package entity

type Filter struct {
	limit int
	page  int
}

type FilterDTO struct {
	Limit int
	Page  int
}

func NewFilter(dto *FilterDTO) *Filter {
	if dto.Limit <= 0 {
		dto.Limit = 10
	}
	if dto.Limit > 100 {
		dto.Limit = 100
	}
	if dto.Page < 0 {
		dto.Page = 1
	}
	return &Filter{
		limit: dto.Limit,
		page:  dto.Page,
	}
}
