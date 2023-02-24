package models

type AccountModel struct {
	Email    string `dbq:"email"`
	Nama     string `dbq:"nama"`
	Password string `dbq:"password"`
}

func GetTableName() string {
	return "account"
}

func TableAccount() []string {
	return []string{
		"email",
		"nama",
		"password",
	}
}
