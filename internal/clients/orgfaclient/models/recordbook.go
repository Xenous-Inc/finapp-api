package models

type RecordBookItem struct {
	Semester    []RecordBookSemesterItem `json:"semesters"`
	AverageMark float32                      `json:"average_mark"`
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
	Scale        int    `json:"scale"`
}
