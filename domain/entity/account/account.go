package account

import (
	"CareerCenter/domain/valueobject"
	"CareerCenter/utils"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	id         string
	email      string
	name       string
	password   string
	role       valueobject.TypeRoles
	createdAt  time.Time
	updatedAt  time.Time
	deletedAt  time.Time
	registerBy string
}

type AccountDTO struct {
	Id         string
	Email      string
	Name       string
	Password   string
	Role       valueobject.TypeRoles
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	RegisterBy string
}

func NewAccount(dto *AccountDTO) (*Account, error) {
	err := dto.Validation()
	if err != nil {
		return nil, err
	}
	return &Account{
		id:         uuid.NewString(),
		email:      dto.Email,
		name:       dto.Name,
		password:   dto.Password,
		role:       dto.Role,
		createdAt:  dto.CreatedAt,
		updatedAt:  dto.UpdatedAt,
		deletedAt:  dto.DeletedAt,
		registerBy: dto.RegisterBy,
	}, nil
}

func (dto *AccountDTO) Validation() error {
	email := utils.ValitEmail(dto.Email)
	if email != true {
		return errors.New("error format email")
	}

	timeNow := time.Now()
	if dto.CreatedAt.IsZero() && dto.UpdatedAt.IsZero() {
		dto.CreatedAt = timeNow
		dto.UpdatedAt = timeNow
	}

	if dto.Role.StringRoles() == "" {
		dto.Role = valueobject.NewTypeRoles(valueobject.ROLE_MEMBER)
	}
	return nil
}

func (g *Account) GetId() string {
	return g.id
}
func (g *Account) GetEmail() string {
	return g.email
}
func (g *Account) GetName() string {
	return g.name
}
func (g *Account) GetPassword() string {
	return g.password
}
func (g *Account) GetRole() valueobject.TypeRoles {
	return g.role
}
func (g *Account) GetCreatedAt() time.Time {
	return g.createdAt
}
func (g *Account) GetUpdatedAt() time.Time {
	return g.updatedAt
}
func (g *Account) GetDeletedAt() time.Time {
	return g.deletedAt
}
func (g *Account) GetRegisterBy() string {
	return g.registerBy
}

func (g *Account) SetRegisterBy(email string) {
	g.registerBy = email
}

type Login struct {
	Token string
	Role  string
}

func (g *Login) SetLogin(token string, role string) {
	g.Token = token
	g.Role = role
}
