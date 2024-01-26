package dto

import "github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"

type MyGroup struct {
	Id       int          `json:"id"`
	UserId   int          `json:"user_id"`
	FullName string       `json:"fullname"`
	Photo    string `json:"photo"`
}

func MyGroupFromClientModel(m *models.Student) MyGroup {
	return MyGroup{
		Id:       m.Id,
		UserId:   m.UserId,
		FullName: m.FullName,
		Photo:    m.Photo.Original,
	}
}
