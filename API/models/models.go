package models

type Search struct {
	Search_type string   `json:"search_type"`
	Query       QueryObj `json:"query"`
	Sort_fields []string `json:"sort_fields"`
	From        int      `json:"from"`
	Max_results int      `json:"max_results"`
	Source      []string `json:"_source"`
}

type QueryObj struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}

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
