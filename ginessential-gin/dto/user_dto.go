package dto

import "ginessential/model"

type UserDto struct {
	Name string `json:"name"`
	Tel  string `json:"tel"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		user.Name,
		user.Tel,
	}
}
