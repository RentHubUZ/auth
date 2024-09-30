package models

type RegisterData struct { // for swagger
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RegisterReq struct {
	FullName       string
	Email          string
	PhoneNumber    string
	HashedPassword string
}

type RegisterResp struct{ ID string }

type UserDetails struct {
	ID             string
	HashedPassword string
	Role           string
}
