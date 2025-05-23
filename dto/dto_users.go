package dto

type UserDto struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UsersDto []UserDto