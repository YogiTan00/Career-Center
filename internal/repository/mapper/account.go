package mapper

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/models"
	"github.com/rocketlaunchr/dbq/v2"
)

func EntityToModel(m *entity.Account) *models.AccountModel {
	return &models.AccountModel{
		Email:    m.GetEmail(),
		Nama:     m.GetNama(),
		Password: m.GetPassword(),
	}
}

func ModelToEntity(m *models.AccountModel) (*entity.AccountDTO, error) {
	data := &entity.AccountDTO{
		Email:    m.Email,
		Nama:     m.Nama,
		Password: m.Password,
	}
	return data, nil
}

func EntityToInterface(data *entity.Account) []interface{} {
	return dbq.Struct(EntityToModel(data))
}

func DomainToInterface(domain *entity.Account) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, EntityToInterface(domain))
	return
}
