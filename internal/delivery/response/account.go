package response

type LoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
