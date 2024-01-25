package models

import "github.com/Xenous-Inc/finapp-api/internal/utils/jsoner"

type Teacher struct {
	Id          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

func (t *Teacher) String() string {
	return jsoner.Jsonify(t)
}
