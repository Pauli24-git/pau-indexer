package models

type Credentials struct {
	Id       string `json:"_id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Validated bool `json:"validated"`
	User      UserData
}

type UserData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
