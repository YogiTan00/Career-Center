package request

type RequestRegister struct {
	Email    string `json:"email"`
	Nama     string `json:"name"`
	Password string `json:"password"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
