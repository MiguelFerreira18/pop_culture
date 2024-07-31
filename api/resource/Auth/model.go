package auth

type FormLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Message string
	Token   string
}

func NewLoginDTO(message string, token string) *LoginDTO {
	return &LoginDTO{
		Message: message,
		Token:   token,
	}

}
