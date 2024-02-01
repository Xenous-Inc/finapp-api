package models

type Profile struct {
	Id       int            `json:"id"`
	Type     string         `json:"type"`
	MiniUser MiniUser       `json:"user"`
	Faculty  Faculty        `json:"faculty"`
	EduGroup EducationGroup `json:"edu_group"`
}

type ProfileDetails struct {
	Id               int                    `json:"id"`
	UserId           int                    `json:"user_id"`
	Type             string                 `json:"type"`
	EduForm          string                 `json:"edu_form"`
	EduMarkBookNum   string                 `json:"edu_mark_book_num"`
	FacultyId        int                    `json:"faculty_id"`
	EduGroupId       int                    `json:"edu_group_id"`
	EduCourse        int                    `json:"edu_course"`
	EduSemester      int                    `json:"edu_semester"`
	EduYear          int                    `json:"edu_year"`
	Role             string                 `json:"role"`
	TypeName         string                 `json:"type_name"`
	BitrixLogin      string                 `json:"bitrix_login"`
	BitrixEmail      string                 `json:"bitrix_email"`
	User             User                   `json:"user"`
	Faculty          Faculty                `json:"faculty"`
	EduGroup         EducationGroup         `json:"edu_group"`
	EduQualification EducationQualification `json:"edu_qualification"`
}

type MiniUser struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	FullName string `json:"fullname"`
	LastName string `json:"lastname"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Photo    Photo  `json:"photo"`
}

type User struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	BitrixId  int    `json:"bitrix_id"`
	FullName  string `json:"fullname"`
	LastName  string `json:"lastname"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Sex       string `json:"sex"`
	Birthdate string `json:"birthdate"`
	Phone     string `json:"phone"`
	Photo     Photo  `json:"photo"`
}

type Faculty struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type EducationGroup struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type EducationQualification struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
