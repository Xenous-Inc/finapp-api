package dto

import "github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"

type Order struct {
	Id          int    `json:"id" example:"192754"`
	Title       string `json:"title" example:"О переводе обучающихся на следующий курс (безусловно, условно)"`
	Date        string `json:"date" example:"2023.09.01"`
	ApproveDate string `json:"approveDate" example:"2023.09.01"`
	Number      string `json:"number" example:"3569/у-с"`
} //@name Order

func OrderFromClientModel(m *models.Order) Order {
	return Order{
		Id:          m.Id,
		Title:       m.Titles,
		Date:        m.Date,
		ApproveDate: m.DateApprove,
		Number:      m.Number,
	}
}
