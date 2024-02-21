package models

import "github.com/Xenous-Inc/finapp-api/internal/utils/jsoner"

type Teacher struct {
	Id          int      `json:"id"`
	UserId      int      `json:"user_id"`
	FullName    string   `json:"fullname"`
	Photo       Photo    `json:"photo"`
	Disciplines []string `json:"disciplines"`
}

func (s *Teacher) String() string {
	return jsoner.Jsonify(s)
}
