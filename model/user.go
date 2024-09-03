package model

type UserResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UserRequest struct {
	Id      []int  `form:"id"`
	Name    string `form:"name"`
	Address string `form:"address"`
}

type BulkDelete struct {
	Id []int `json:"id"`
}
