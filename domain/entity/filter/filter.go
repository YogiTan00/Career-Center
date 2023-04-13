package filter

type Filter struct {
	q      string
	limit  uint32
	page   uint32
	status bool
}

type FilterDTO struct {
	Q      string
	Limit  uint32
	Page   uint32
	Status bool
}

func NewFilter(dto *FilterDTO) *Filter {
	if dto.Limit > 100 {
		dto.Limit = 100
	}
	if dto.Page == 0 {
		dto.Page = 1
	}
	page := (dto.Page - 1) * dto.Limit
	return &Filter{
		q:      dto.Q,
		limit:  dto.Limit,
		page:   page,
		status: dto.Status,
	}
}

func (data *Filter) GetQ() string {
	return data.q
}

func (data *Filter) GetLimit() uint32 {
	return data.limit
}

func (data *Filter) GetPage() uint32 {
	return data.page
}

func (data *Filter) GetStatus() bool {
	return data.status
}
