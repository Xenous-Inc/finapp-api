package dto

import "github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"

type RecordBookItem struct {
	Semester []RecordBookSemesterItem `json:"semesters"`
}

type RecordBookSemesterItem struct {
	SemesterNumber int                          `json:"semester"`
	Data           []RecordBookSemesterItemData `json:"data"`
}

type RecordBookSemesterItemData struct {
	Date         string `json:"date"`
	Mark         int    `json:"mark"`
	Subject      string `json:"subject"`
	LecturerName string `json:"lecturers"`
	T1           int    `json:"t1"`
	T2           int    `json:"t2"`
	Uo           int    `json:"uo"`
	Result       int    `json:"itog"`
}

func RecordBookItemFromClientModel(m []models.RecordBookItem) []RecordBookItem {
	var studyPlans []RecordBookItem
	for _, stud := range m {
		studyPlans = append(studyPlans, RecordBookItem{
			Semester: func() []RecordBookSemesterItem {
				var semesters []RecordBookSemesterItem
				for _, semester := range stud.Semester {
					data := make([]RecordBookSemesterItemData, 0)
					for _, item := range semester.Data {
						data = append(data, RecordBookSemesterItemData{
							Date:         item.Date,
							Mark:         item.Mark,
							Subject:      item.Subject,
							LecturerName: item.LecturerName,
							T1:           item.T1,
							T2:           item.T2,
							Uo:           item.Uo,
							Result:       item.Result,
						})
					}
					semesters = append(semesters, RecordBookSemesterItem{
						SemesterNumber: semester.SemesterNumber,
						Data:           data,
					})
				}
				return semesters
			}(),
		})
	}
	return studyPlans
}
