package dto

import "github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"

type StudentCard struct {
	URL string `json:"url" example:"https://org.fa.ru/upload/temp/49e6475a8dedba8800cbb8170f5fab30.pdf"`
} //@name StudentCard

func StudentCardFromClientModel(m *models.StudentCard) StudentCard {
	return StudentCard{
		URL: models.BASE_URL + m.Path,
	}
}
