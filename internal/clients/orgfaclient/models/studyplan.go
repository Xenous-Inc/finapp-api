package models

type StudyPlan struct {
	Title    string              `json:"title"`
	Semester []StudyPlanSemester `json:"semesters"`
}

type StudyPlanSemester struct {
	Semester int       `json:"semester"`
	Section  []Section `json:"sections"`
}

type Section struct {
	Title string `json:"title"`
	Hours int    `json:"hours"`
	Terms []Term `json:"terms"`
}

// TODO: init me
type SubTerm struct {
}

// TODO: init me
type Subject struct {
}

type Term struct {
	ExternalId string `json:"external_id"`
	Number     int    `json:"num"`
	Exam       int    `json:"exam"`
	Test       int    `json:"test"`
	Kp         int    `json:"kp"`
	Kr         int    `json:"kr"`
}
