package mapper

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/domain/valueobject"
	"CareerCenter/internal/repository/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func EntityToModel(m *account.Account) *models.AccountModel {
	return &models.AccountModel{
		Id:        m.GetId(),
		Email:     m.GetEmail(),
		Nama:      m.GetName(),
		Password:  m.GetPassword(),
		Role:      m.GetRole().StringRoles(),
		CreatedAt: m.GetCreatedAt(),
		UpdateAt:  m.GetUpdatedAt(),
	}
}

func ModelToEntity(m *models.AccountModel) *account.AccountDTO {
	role := valueobject.NewTypeRolesFromString(m.Role)
	data := &account.AccountDTO{
		Email:      m.Email,
		Name:       m.Nama,
		Password:   m.Password,
		Role:       role,
		CodeOtp:    m.CodeOtp,
		ExpiredOtp: m.ExpiredOtp,
	}
	return data
}

func EntityToInterface(data *account.Account) []interface{} {
	return dbq.Struct(EntityToModel(data))
}

func DomainToInterface(domain *account.Account) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, EntityToInterface(domain))
	return
}
