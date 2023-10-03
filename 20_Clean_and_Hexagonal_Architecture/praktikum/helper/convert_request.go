package helper

import (
	"simple_clean_code/model/domain"
	"simple_clean_code/model/web"
)

func UserCreateRequestToUserDomain(request web.UserCreateRequest) *domain.User {
	return &domain.User{
		Email: request.Email,
		Password: request.Password,
	}
}