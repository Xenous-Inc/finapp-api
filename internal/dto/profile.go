package dto

type Profile struct {
	Id           string `json:"id" example:"93492"`
	TicketNumber string `json:"ticket" example:"226299"`
	Name         string `json:"name" example:"Дредноут Александр Дмитриевич"`
	Group        string `json:"group" example:"ФФ22-4"`
	FacultyName  string `json:"faculty" example:"Финансовый факультет"`
	ImageURL     string `json:"image" example:"https://org.fa.ru/bitrix/galaktika/galaktika.vuzapi/public/files/users/89105/ctn-25001.281474976824101_optimized.jpg"`
} //@name Profile
// TODO: Func create from Client Model

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
} //@name ProfileDetails
