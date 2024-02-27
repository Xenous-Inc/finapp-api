package dto

import "github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"

type Student struct {
	Id       int    `json:"id" example:"87436"`
	FullName string `json:"name" example:"Шереметьев Александр Дмитриевич"`
	Cover    string `json:"cover" example:"https://org.fa.ru/bitrix/galaktika/galaktika.vuzapi/public/files/users/83066/25001.281474976821736_optimized.jpg"`
} //@name StudentItem

func StudentFromClientModel(m *models.Student) Student {
	cover := ""
	if m.Photo.Original != "" {
		cover = models.BASE_URL + m.Photo.Original
	}

	return Student{
		Id:       m.Id,
		FullName: m.FullName,
		Cover:    cover,
	}
}
