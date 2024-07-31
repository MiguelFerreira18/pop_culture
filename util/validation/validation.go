package validation

import (
	"errors"
	"pop_culture/util/hash"
	"regexp"
	"strconv"
	"unicode"
)

// User Rules
func UserNameRules(name string) (*string, error) {
	pattern := `^[a-zA-Z0-9 ]+$`
	if len(name) < 4 || len(name) > 20 {
		return nil, errors.New("Name is not the correct size: " + strconv.Itoa(len(name)))
	}
	if !matchPattern(name, pattern) {
		return nil, errors.New("name is not conformed")
	}
	return &name, nil
}

func EmailRules(email *string) (*string, error) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if matchPattern(*email, pattern) {
		return email, nil
	}

	return nil, errors.New("email does not follow the rules")

}
func PasswordRules(password string) (*string, error) {
	if !passwordPattern(password) {
		return nil, errors.New("User password is not conformed")
	}
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return nil, err
	}
	return &hashedPassword, nil
}

// Media Type Rules
func MediaNameRules(media string) (*string, error) {
	pattern := `(?i)(\b(select|union|insert|update|delete|drop|alter|create|shutdown|exec)\b|\-\-|\;|\#|\')`
	if len(media) < 3 || len(media) > 15 {
		return nil, errors.New("Media name is not the correct size: " + strconv.Itoa(len(media)))
	}
	if !matchPattern(media, pattern) {
		return nil, errors.New("media name is not conformed")
	}

	return &media, nil
}

// Attribute Type Rules
func AttributeNameRules(name string) (*string, error) {
	pattern := `^[a-zA-Z0-9 ]+$`
	if len(name) < 3 || len(name) > 15 {
		return nil, errors.New("Name is not the correct size: " + strconv.Itoa(len(name)))
	}
	if !matchPattern(name, pattern) {
		return nil, errors.New("media name is not conformed")
	}
	return &name, nil
}

func RoleNameRules(name string) (*string, error) {
	pattern := `^[a-zA-Z0-9 ]+$`
	if len(name) < 3 || len(name) > 6 {
		return nil, errors.New("Name of the Role is not of the correct size: " + strconv.Itoa(len(name)))
	}
	if !matchPattern(name, pattern) {
		return nil, errors.New("Role Name is not conformed to the rules")
	}
	return &name, nil
}
func RoleDescriptionRules(description string) (*string, error) {
	pattern := `^[a-zA-Z0-9 ]+$`
	if len(description) < 8 || len(description) > 256 {
		return nil, errors.New("Description of the Role is not of the correct size: " + strconv.Itoa(len(description)))
	}
	if !matchPattern(description, pattern) {
		return nil, errors.New("Role description is not conformed to the rules")
	}
	return &description, nil
}

func matchPattern(input string, pattern string) bool {

	re := regexp.MustCompile(pattern)

	return re.MatchString(input)

}

func passwordPattern(password string) bool {
	var (
		hasDigit       bool
		hasSpecialChar bool
		hasUppercase   bool
		hasLowercase   bool
	)
	for _, char := range password {
		switch {
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		}
	}
	return hasLowercase && hasUppercase && hasSpecialChar && hasDigit
}
