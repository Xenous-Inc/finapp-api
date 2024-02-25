package dto

import "github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"

type RecordBookItem struct {
	Semester []RecordBookSemesterItem `json:"semesters"`
} //@name RecordBookItem

type RecordBookSemesterItem struct {
	SemesterNumber int                          `json:"semester"`
	Data           []RecordBookSemesterItemData `json:"data"`
} //@name RecordBookSemesterItem

type RecordBookSemesterItemData struct {
	Date           string `json:"date"`
	Mark           int    `json:"mark"`
	Subject        string `json:"subject"`
	LecturerName   string `json:"lecturers"`
	CurrentControl int    `json:"current control"`
	WorkInSemester int    `json:"work_in_semester"`
	ExamOrTest     int    `json:"exam_or_test"`
	Result         int    `json:"itog"`
	Scale          int    `json:"scale"`
} //@name RecordBookSemesterItemData

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
							Date:           item.Date,
							Mark:           item.Mark,
							Subject:        item.Subject,
							LecturerName:   item.LecturerName,
							CurrentControl: item.T1,
							WorkInSemester: item.T2,
							ExamOrTest:     item.Uo,
							Result:         item.Result,
							Scale:          item.Scale,
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
