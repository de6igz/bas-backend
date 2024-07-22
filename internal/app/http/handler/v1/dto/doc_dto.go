package dto

import "bas-backend/domain/model"

type DocDto struct {
	Items []model.Document `json:"docs"`
}
