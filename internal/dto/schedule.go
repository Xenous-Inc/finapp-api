package dto

type GetScheduleRequest struct {
	EntityId   string `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}