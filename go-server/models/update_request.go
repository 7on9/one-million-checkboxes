package models

type UpdateRequest struct {
	Value    bool `json:"value"`
	Position int  `json:"position"`
}
