package validation

import (
	"errors"
	"pop_culture/util/hash"
	"regexp"
	"strconv"
)

// User Rules
func UserNameRules(name string) (*string, error) {
	pattern := `(?i)(\b(select|union|insert|update|delete|drop|alter|create|shutdown|exec)\b|\-\-|\;|\#|\')`
	if len(name) < 5 && len(name) > 20 {
		return nil, errors.New("Name is not the correct size: " + strconv.Itoa(len(name)))
	}
	if matchPattern(name, pattern) {
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
	pattern := `(?i)(\b(select|union|insert|update|delete|drop|alter|create|shutdown|exec)\b|\-\-|\;|\#|\')`

	if matchPattern(password, pattern) {
		return nil, errors.New("the password does not follow the rules")
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
	if len(media) < 3 && len(media) > 15 {
		return nil, errors.New("Media name is not the correct size: " + strconv.Itoa(len(media)))
	}
	if matchPattern(media, pattern) {
		return nil, errors.New("media name is not conformed")
	}

	return &media, nil
}

// Attribute Type Rules
func AttributeNameRules(name string) (*string, error) {
	pattern := `(?i)(\b(select|union|insert|update|delete|drop|alter|create|shutdown|exec)\b|\-\-|\;|\#|\')`
	if len(name) < 3 && len(name) > 15 {
		return nil, errors.New("Name is not the correct size: " + strconv.Itoa(len(name)))
	}
	if matchPattern(name, pattern) {
		return nil, errors.New("media name is not conformed")
	}
	return &name, nil
}

func matchPattern(input string, pattern string) bool {

	re := regexp.MustCompile(pattern)

	return re.MatchString(input)

}
