package dto

type Order struct {
	Id          int    `json:"id"`
	ExternalId  string `json:"external_id"`
	Number      string `json:"number"`
	Date        string `json:"date"`
	DateApprove string `json:"date_approve"`
	Titles      string `json:"title"`
	Action      string `json:"action"`
	Pivot       Pivot  `json:"pivot"`
}

type Pivot struct {
	ProfileId int `json:"profile_id"`
	OrderId   int `json:"order_id"`
}
