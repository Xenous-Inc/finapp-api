package dto

import "github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"

type Profile struct {
	Id           int    `json:"id" example:"93492"`
	TicketNumber string `json:"ticket" example:"226299"`
	Name         string `json:"name" example:"Дредноут Александр Дмитриевич"`
	Group        string `json:"group" example:"ФФ22-4"`
	FacultyName  string `json:"faculty" example:"Финансовый факультет"`
	ImageURL     string `json:"image" example:"https://org.fa.ru/bitrix/galaktika/galaktika.vuzapi/public/files/users/89105/ctn-25001.281474976824101_optimized.jpg"`
} //@name Profile

func ProfileFromClientModel(m *models.Profile) Profile {
	return Profile{
		Id:           m.Id,
		TicketNumber: m.MiniUser.Login,
		Name:         m.MiniUser.FullName,
		Group:        m.EduGroup.Title,
		FacultyName:  m.Faculty.Title,
		ImageURL:     models.BASE_URL + m.MiniUser.Photo.Thumbnail,
	}
}

type ProfileDetails struct {
	Id                 int    `json:"id" example:"93492"`
	TicketNumber       string `json:"ticket" example:"226299"`
	Name               string `json:"name" example:"Дредноут Александр Дмитриевич"`
	Group              string `json:"group" example:"ФФ22-4"`
	FacultyName        string `json:"faculty" example:"Финансовый факультет"`
	EducationDirection string `json:"direction" example:"Экономика"`
	EducationForm      string `json:"form" example:"Очная"`
	EducationLevel     string `json:"level" example:"Бакалавр"`
	CourseNumber       uint8  `json:"course" example:"2"`
	SemesterNumber     uint8  `json:"semester" example:"3"`
	EduGroup           string `json:"group_id"`
	ImageURL           string `json:"image" example:"https://org.fa.ru/bitrix/galaktika/galaktika.vuzapi/public/files/users/89105/ctn-25001.281474976824101_optimized.jpg"`
} //@name ProfileDetails

func ProfileDetailsFromClientModel(m *models.ProfileDetails) ProfileDetails {
	return ProfileDetails{
		Id:                 m.Id,
		TicketNumber:       m.User.Login,
		Name:               m.User.FullName,
		Group:              m.EduGroup.Title,
		FacultyName:        m.Faculty.Title,
		EducationDirection: m.EduQualification.Title,
		EducationForm:      m.EduForm,
		EducationLevel:     m.EduMarkBookNum,
		CourseNumber:       uint8(m.EduCourse),
		SemesterNumber:     uint8(m.EduSemester),
		ImageURL:           models.BASE_URL + m.User.Photo.Thumbnail,
	}
}
