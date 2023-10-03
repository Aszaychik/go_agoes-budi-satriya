package helper

import (
	"simple_clean_code/model/domain"
	"simple_clean_code/model/web"
)

func UserDomainToUserCreateResponse(user *domain.User) web.UserCreateResponse {
	return web.UserCreateResponse{
		Email: user.Email,
	}
}