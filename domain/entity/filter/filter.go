package filter

type Filter struct {
	limit uint32
	page  uint32
}

type FilterDTO struct {
	Limit uint32
	Page  uint32
}

func NewFilter(dto *FilterDTO) *Filter {
	if dto.Limit > 100 {
		dto.Limit = 100
	}
	if dto.Page < 0 {
		dto.Page = 1
	}
	page := (dto.Page - 1) * dto.Limit
	return &Filter{
		limit: dto.Limit,
		page:  page,
	}
}

func (dto *Filter) GetLimit() uint32 {
	return dto.limit
}

func (dto *Filter) GetPage() uint32 {
	return dto.page
}
