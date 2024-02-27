package models

import "github.com/Xenous-Inc/finapp-api/internal/utils/jsoner"

type Student struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Photo    Photo  `json:"photo"`
}

type Photo struct {
	Original  string `json:"orig"`
	Thumbnail string `json:"thumbnail"`
	Small     string `json:"small"`
}

func (s *Student) String() string {
	return jsoner.Jsonify(s)
}
