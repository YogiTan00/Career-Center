package account

import (
	"CareerCenter/internal/config/database"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mysSqlConn  = database.InitMysqlDB()
	ctx         = context.TODO()
	repoAccount = NewAccountMysqlInteractor(mysSqlConn)
)

func TestAccountMysqlInteractor_GetByEmail(t *testing.T) {
	transactionDetail, err := repoAccount.GetByEmail(ctx, "jIzijthBYx@gmail.com")
	//check := utils.CheckPasswordHash("abcd123", transactionDetail.Password)
	//fmt.Println(check)
	fmt.Println(transactionDetail)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}
