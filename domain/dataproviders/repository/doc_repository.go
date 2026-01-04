package repository

import (
	"bas-backend/domain/model"
	"context"
	"database/sql"
	"log"
)

type DocRepository interface {
	// GetAllDocs Получить все документы
	GetAllDocs(ctx context.Context) ([]model.Document, error)
}

type docRepository struct {
	db *sql.DB
}

func NewDocRepository(ctx context.Context, db *sql.DB) DocRepository {
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	return &docRepository{
		db: db,
	}
}

// GetAllDocs Получить все документы
func (r *docRepository) GetAllDocs(ctx context.Context) ([]model.Document, error) {
	sqlText := `select preview_url, document_url from docs`
	rows, err := r.db.QueryContext(ctx, sqlText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	docs := make([]model.Document, 0)
	for rows.Next() {
		var doc model.Document
		if err := rows.Scan(&doc.PreviewURL, &doc.DocumentURL); err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, rows.Err()
}
