package filter

type Filter struct {
	q      string
	page   uint32
	limit  uint32
	status bool
	order  string
}

type FilterDTO struct {
	Q      string
	Page   uint32
	Limit  uint32
	Status bool
	Order  string
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
		page:   page,
		limit:  dto.Limit,
		status: dto.Status,
		order:  dto.Order,
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

func (data *Filter) GetOrder() string {
	return data.order
}
