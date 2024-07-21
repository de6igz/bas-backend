package repository

import (
	"bas-backend/config"
	"bas-backend/domain/model"
	"context"
	"github.com/go-pg/pg/v10"
	"log"
	"strconv"
)

type PartnerRepository interface {
	GetAll(ctx context.Context) ([]model.Partner, error)
}

type partnerRepository struct {
	db *pg.DB
}

func NewPartnerRepository(ctx context.Context, config *config.Config) PartnerRepository {

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

	return &partnerRepository{
		db: connection,
	}
}

func (p *partnerRepository) GetAll(ctx context.Context) ([]model.Partner, error) {
	sql := "select url,description from partners"
	var partners []model.Partner
	_, err := p.db.QueryContext(ctx, &partners, sql)
	if err != nil {
		log.Printf("error getting partners: %v", err)

		return nil, err
	}

	return partners, nil
}
