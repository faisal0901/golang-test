package Response

type LoginResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"Password"`
	Token    string `json:"token"`
}