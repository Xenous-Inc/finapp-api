package dto

type Teacher struct {
	Id          string `json:"id" example:"00000000-0001-2345-6789-000000005451"`
	Title       string `json:"title" example:"Махаматов Таир"`
	Description string `json:"desc" example:"Департамент гуманитарных наук"`
} // @name Teacher
