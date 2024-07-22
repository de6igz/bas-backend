package model

import "context"

type Document struct {
	URL string `json:"url", pg:"url"`
}

type DocRepository interface {
	GetAllDocs(ctx context.Context) ([]Document, error)
}
