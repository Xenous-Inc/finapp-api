package dto

import "github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"

type Teacher struct {
	Id          string `json:"id" example:"00000000-0001-2345-6789-000000005451"`
	Title       string `json:"title" example:"Махаматов Таир"`
	Description string `json:"desc" example:"Департамент гуманитарных наук"`
} // @name Teacher


type TeacherGroup struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userId"`
	FullName string `json:"name"`
	Cover    string `json:"cover"`
	Disciplines []string `json:"disciplines"`
} // @name TeacherGroup

func TeacherGroupFromClientModel(m *models.Teacher) TeacherGroup {
	cover := ""
	if m.Photo.Original != "" {
		cover = models.BASE_URL + m.Photo.Original
	}

	return TeacherGroup{
		Id:       m.Id,
		UserId:   m.UserId,
		FullName: m.FullName,
		Cover:    cover,
		Disciplines: m.Disciplines,
	}
}