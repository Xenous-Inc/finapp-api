package dto

type StudyPlan struct {
	Id         int                 `json:"id"`
	ExternalId string              `json:"external_id"`
	Title      string              `json:"title"`
	Number     string              `json:"number"`
	EduYear    int                 `json:"edu_year"`
	Okso       string              `json:"okso"`
	Year       string              `json:"year"`
	FacultyId  int                 `json:"faculty_id"`
	EduCource  int                 `json:"edu_cource"`
	EisId      string              `json:"eis_id"`
	Semester   []StudyPlanSemester `json:"semesters"`
}

type StudyPlanSemester struct {
	Semester int       `json:"semester"`
	Section  []Section `json:"sections"`
}

type Section struct {
	Id         int    `json:"id"`
	ExternalId string `json:"external_id"`
	RupId      int    `json:"rup_id"`
	Title      string `json:"title"`
	Type       string `json:"section"`
	Hours      int    `json:"hours"`
	Zet        string `json:"zet"`
	// SubTerms   []SubTerm `json:"sub_terms"`
	// Subjects   []Subject `json:"subjects"`
	Terms []Term `json:"terms"`
}

// TODO: init me
type SubTerm struct {
}

// TODO: init me
type Subject struct {
}

type Term struct {
	Id           int    `json:"id"`
	ExternalId   string `json:"external_id"`
	RupSectionId int    `json:"rup_section_id"`
	Number       int    `json:"num"`
	Exam         int    `json:"exam"`
	Test         int    `json:"test"`
	Kp           int    `json:"kp"`
	Kr           int    `json:"kr"`
	TestWork     int    `json:"testwork"`
	Labs         int    `json:"labs"`
	Lections     int    `json:"lections"`
	Practice     int    `json:"practice"`
	Self         int    `json:"self"`
}
