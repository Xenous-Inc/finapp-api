package dto

import "time"

type Date string

func (d *Date) UnmarshalJSON(bytes []byte) error {
	date, err := time.Parse(`"2006-01-02T15:04:05.000-0700"`, string(bytes))
	if err != nil {
		return err
	}
	*d = Date(date.Format("2024.12.01"))

	return nil
}

type GetScheduleRequest struct {
	EntityId  string `json:"id" validate:"required"`
	StartDate Date   `json:"startDate"`
	EndDate   Date   `json:"endDate"`
}

