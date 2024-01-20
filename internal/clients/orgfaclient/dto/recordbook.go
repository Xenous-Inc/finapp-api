package dto

type RecordBookItem struct {
	Year     int                      `json:"year"`
	Semester []RecordBookSemesterItem `json:"semesters"`
}

type RecordBookSemesterItem struct {
	SemesterNumber int                          `json:"semester"`
	Data           []RecordBookSemesterItemData `json:"data"`
}

type RecordBookSemesterItemData struct {
	SemesterNumber int    `json:"semester"`
	Date           string `json:"date"`
	Year           int    `json:"year"`
	HoursCount     string `json:"hours"`
	ControlType    string `json:"control_type"`
	Mark           int    `json:"mark"`
	MarkTitle      string `json:"mark_title"`
	Subject        string `json:"subject"`
	// Zet string `json:"zet"`
	LecturerName string `json:"lecturers"`
	// //P1 string `json:"p1"`
	// //P2 string `json:"p2"`
	// //T1 string `json:"t1"`
	// //T2 string `json:"t2"`
	// //Uo string `json:"uo"`
	Result int `json:"itog"`
}
