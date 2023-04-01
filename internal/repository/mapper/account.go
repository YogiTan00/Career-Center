package mapper

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/internal/repository/models"
	"github.com/rocketlaunchr/dbq/v2"
)

func EntityToModel(m *account.Account) *models.AccountModel {
	return &models.AccountModel{
		Id:        m.GetId(),
		Email:     m.GetEmail(),
		Nama:      m.GetNama(),
		Password:  m.GetPassword(),
		CreatedAt: m.GetCreatedAt(),
		UpdateAt:  m.GetUpdatedAt(),
	}
}

func ModelToEntity(m *models.AccountModel) (*account.AccountDTO, error) {
	data := &account.AccountDTO{
		Email:    m.Email,
		Name:     m.Nama,
		Password: m.Password,
	}
	return data, nil
}

func EntityToInterface(data *account.Account) []interface{} {
	return dbq.Struct(EntityToModel(data))
}

func DomainToInterface(domain *account.Account) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, EntityToInterface(domain))
	return
}
