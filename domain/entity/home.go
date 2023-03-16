package entity

type HomePage struct {
	account *Account
	jobs    []*Jobs
}

type HomePageDTO struct {
	account *Account
	Jobs    []*JobsDTO
}

func NewHomePage(dto *HomePageDTO) (*HomePage, error) {
	listJob := make([]*Jobs, 0)
	for _, data := range dto.Jobs {
		jobs := &JobsDTO{
			Id:        data.Id,
			Position:  data.Position,
			Company:   data.Company,
			Logo:      data.Logo,
			Address:   data.Address,
			CreatedAt: data.CreatedAt,
		}
		resultJob, err := NewJobs(jobs)
		if err != nil {
			return nil, err
		}
		listJob = append(listJob, resultJob)
	}

	return &HomePage{
		account: &Account{},
		jobs:    listJob,
	}, nil
}

func (dto *HomePageDTO) Validation() error {
	return nil
}
