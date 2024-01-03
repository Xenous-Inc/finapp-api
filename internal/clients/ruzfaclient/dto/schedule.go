package dto

import "github.com/Xenous-Inc/finapp-api/internal/utils/jsoner"

type Schedule struct {
	//
	Auditorium    string `json:"auditorium"`
	AuditoriumOid int    `json:"auditoriumOid"`

	BeginLesson string `json:"beginLesson"`
	EndLesson   string `json:"endLesson"`

	// Street/address
	Building    string `json:"building"`
	BuildingOid int    `json:"buildingOid"`

	// lesson
	Discipline    string `json:"discipline"`
	DisciplineOid int    `json:"disciplineOid"`

	// date lesson
	Date string `json:"date"`
	//DayOfWeek string `json:"dayOfWeek"`
	DayOfWeekString string `json:"dayOfWeekString"`

	// type lesson
	KindOfWork    string `json:"kindOfWork"`
	KindOfWork0id int    `json:"kindOfWorkOid"`

	// techear
	Lecturer    string `json:"lecturer"`
	Lecturer0id int    `json:"lecturerOid"`
}

func (s *Schedule) String() string {
	return jsoner.Jsonify(s)
}
