package repository

import (
	"CareerCenter/internal/config/database"
	"CareerCenter/utils"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mysSqlConn  = database.InitMysqlDB()
	ctx         = context.TODO()
	repoAccount = NewAccountMysqlInteractor(mysSqlConn)
)

func TestAccountMysqlInteractor_GetByEmail(t *testing.T) {
	transactionDetail, err := repoAccount.GetByEmail(ctx, "yogiwijaya00@gmail.com")
	check := utils.CheckPasswordHash("abcd1234", transactionDetail.Password)
	fmt.Println(transactionDetail)
	fmt.Println(check)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}
