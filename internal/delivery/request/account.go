package request

type RequestRegister struct {
	Email    string `json:"username"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
}
