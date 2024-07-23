package model

import "context"

type Document struct {
	DocumentURL string `json:"document_url" pg:"document_url"`
	PreviewURL  string `json:"preview_url" pg:"preview_url"`
}

type DocRepository interface {
	GetAllDocs(ctx context.Context) ([]Document, error)
}
