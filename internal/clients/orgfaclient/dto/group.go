package dto

type Group struct {
	Id string `json:"id"`
	Label string `json:"label"`
	Description string `json:"description"`
}

func NewGroup(group *Group) *Group {
	return &Group{}
}