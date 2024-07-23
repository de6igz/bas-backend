package repository

import (
	"bas-backend/config"
	"bas-backend/domain/model"
	"context"
	"github.com/go-pg/pg/v10"
	"log"
	"strconv"
)

type DocRepository interface {
	// GetAllDocs Получить все документы
	GetAllDocs(ctx context.Context) ([]model.Document, error)
}

type docRepository struct {
	db *pg.DB
}

func NewDocRepository(ctx context.Context, config *config.Config) DocRepository {

	connection := pg.Connect(&pg.Options{
		Addr:            config.Database.Host + ":" + strconv.Itoa(config.Database.Port),
		User:            config.Database.User,
		Password:        config.Database.Password,
		Database:        config.Database.Name,
		MaxRetries:      3,
		MaxRetryBackoff: 3,
	})

	err := connection.Ping(ctx)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	return &docRepository{
		db: connection,
	}
}

// GetAllDocs Получить все документы
func (r *docRepository) GetAllDocs(ctx context.Context) ([]model.Document, error) {
	sql := `select preview_url,document_url  from docs`
	docs := make([]model.Document, 0)
	_, err := r.db.QueryContext(ctx, &docs, sql)
	if err != nil {

		return docs, err
	}

	return docs, nil
}
