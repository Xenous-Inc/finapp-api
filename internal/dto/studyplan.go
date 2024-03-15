package dto

import "github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"

type StudyPlan struct {
	Title    string              `json:"title"`
	Semester []StudyPlanSemester `json:"semesters"`
} //@name StudyPlan

type StudyPlanSemester struct {
	Semester int       `json:"semester"`
	Section  []Section `json:"sections"`
} //@name StudyPlanSemester

type Section struct {
	Title string `json:"title"`
	Hours int    `json:"hours"`
	Terms []Term `json:"terms"`
} //@name SemesterSection

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
} //@name Term

func StudyPlanFromClientModel(m []models.StudyPlan) []StudyPlan {

	var studyPlans []StudyPlan

	for _, stud := range m {
		studyPlans = append(studyPlans, StudyPlan{
			Title: stud.Title,
			Semester: func() []StudyPlanSemester {
				var semesters []StudyPlanSemester
				for _, semester := range stud.Semester {
					sections := make([]Section, 0)
					for _, section := range semester.Section {
						terms := make([]Term, 0)
						for _, term := range section.Terms {
							terms = append(terms, Term{
								ExternalId: term.ExternalId,
								Number:     term.Number,
								Exam:       term.Exam,
								Test:       term.Test,
								Kp:         term.Kp,
								Kr:         term.Kr,
							})
						}
						sections = append(sections, Section{
							Title: section.Title,
							Hours: section.Hours,
							Terms: terms,
						})
					}
					semesters = append(semesters, StudyPlanSemester{
						Semester: semester.Semester,
						Section:  sections,
					})
				}
				return semesters
			}(),
		})

	}
	return studyPlans
}
