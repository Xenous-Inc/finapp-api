package models

import "github.com/Xenous-Inc/finapp-api/internal/utils/jsoner"

type Auditorium struct {
	Id          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

func (a *Auditorium) String() string {
	return jsoner.Jsonify(a)
}
