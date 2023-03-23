package home

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"context"
)

func (u *UseCaseHomeInteractor) GetHome(ctx context.Context, f *filter.FilterHomeDTO, email string) (*entity.HomePageDTO, error) {
	//filter := filter.NewFilter(f.FilterJobs)
	//
	//find, err := u.repoProfile.GetPhotoProfile(ctx, email)
	//if err != nil {
	//	return nil, err
	//}
	//
	//data, err := u.repoJobs.GetListJobs(ctx, filter)
	//if err != nil {
	//	return nil, err
	//}
	//
	//response := response2.GetHome(data, )
	//return result, nil
	return nil, nil
}
