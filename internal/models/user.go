package models

type User struct {
	ID             int `json:"id_user"`
	Role           string `json:"role"`
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
}
