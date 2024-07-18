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
	GetAll() ([]model.Partner, error)
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

func (p *partnerRepository) GetAll() ([]model.Partner, error) {
	// Здесь нужно добавить логику для работы с базой данных

	return make([]model.Partner, 0), nil
}
