package filter

type FilterHome struct {
	filterJobs     *Filter
	filterActivity *Filter
}

type FilterHomeDTO struct {
	FilterJobs     *FilterDTO
	FilterActivity *FilterDTO
}

func NewFilterHome(dto *FilterHomeDTO) *FilterHome {
	jobs := NewFilter(dto.FilterJobs)
	activity := NewFilter(dto.FilterActivity)
	return &FilterHome{
		filterJobs:     jobs,
		filterActivity: activity,
	}
}
