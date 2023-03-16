package request

type RequestRegister struct {
	Email    string `json:"email"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
