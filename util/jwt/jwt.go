package jwt

import (
	"errors"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type JWT struct {
	UserID uuid.UUID
	Email  *string
	Role   string
}

func NewJWTToken(userID uuid.UUID, email string, role string) *JWT {
	return &JWT{
		UserID: userID,
		Email:  &email,
		Role:   role,
	}
}

func (jwt JWT) Encode(token *jwtauth.JWTAuth) (jwt.Token, *string, error) {
	jwtToken, tokenString, err := token.Encode(map[string]interface{}{"UserID": jwt.UserID, "Email": jwt.Email, "Role": jwt.Role})
	if err != nil {
		return nil, nil, err
	}
	return jwtToken, &tokenString, err
}

func Decode(jwtToken jwt.Token) (*string, *string, *string, error) {

	stringUserID, exists := jwtToken.Get("UserID")
	if !exists {
		return nil, nil, nil, errors.New("There is no user id in the token")
	}
	if err := uuid.Validate(stringUserID.(string)); err != nil {
		return nil, nil, nil, errors.New("The uuid is not valid")
	}
	userID, err := uuid.Parse(stringUserID.(string))
	if err != nil {
		return nil, nil, nil, errors.New("Error parsing the user ID")
	}

	emailInterface, exists := jwtToken.Get("Email")
	if !exists {
		return nil, nil, nil, errors.New("There is no email in the token")
	}
	roleInterface, exists := jwtToken.Get("Role")
	if !exists {
		return nil, nil, nil, errors.New("There is no role in the token")
	}

	email := emailInterface.(string)
	role := roleInterface.(string)
	validUserID := userID.String()

	return &validUserID, &email, &role, nil

}
