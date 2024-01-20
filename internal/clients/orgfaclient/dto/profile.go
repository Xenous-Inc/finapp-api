package dto

type MiniProfile struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	TypeName string `json:"type_name"`
	Name     string `json:"name"`
	Href     string `json:"href"`
	Extra    string `json:"extra"`
}

type AllDataProfile struct {
	Profile []Profile
}

type Profile struct {
	Id               int              `json:"id"`
	UserId           string           `json:"user_id"`
	Type             string           `json:"type"`
	EduForm          string           `json:"edu_form"`
	EduMarkBookNum   string           `json:"edu_mark_book_num"`
	FacultyId        int              `json:"faculty_id"`
	EduGroupId       int              `json:"edu_group_id"`
	EduCourse        int              `json:"edu_course"`
	EduSemester      int              `json:"edu_semester"`
	EduYear          int              `json:"edu_year"`
	Role             string           `json:"role"`
	TypeName         string           `json:"type_name"`
	BitrixLogin      string           `json:"bitrix_login"`
	BitrixEmail      string           `json:"bitrix_email"`
	User             []User           `json:"user"`
	Faculty          Faculty          `json:"faculty"`
	EduGroup         EduGroup         `json:"edu_group"`
	EduQualification EduQualification `json:"edu_qualification"`
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

type EduGroup struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type EduQualification struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
