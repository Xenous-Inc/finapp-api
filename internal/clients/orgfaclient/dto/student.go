package dto

import "github.com/Xenous-Inc/finapp-api/internal/utils/jsoner"

type Student struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	FullName string `json:"fullname"`
	Photo Photo `json:"photo"`
}

type Photo struct {
	Original string	`json:"orig"`
	Thumbnail string `json:"thumbnail"`
	Small string `json:"small"`
}

type Data struct {
	Error int    `json:"error"`
	Student []Student `json:"data"`
}

func (s *Student) String() string {
	return jsoner.Jsonify(s)
}