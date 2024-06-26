package model

type Branch struct {
	Branch_Id int    `json:"branch_id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
}

type IDResp struct {
	Branch_Id int `json:"branch_id"`
}

type BranchMessage struct {
	Message string `json:"message"`
}
