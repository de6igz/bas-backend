package repository

import (
	"bas-backend/domain/model"
	"context"
	"database/sql"
	"log"
)

type PartnerRepository interface {
	GetAllPartners(ctx context.Context) ([]model.Partner, error)
}

type partnerRepository struct {
	db *sql.DB
}

func NewPartnerRepository(ctx context.Context, db *sql.DB) PartnerRepository {
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	return &partnerRepository{
		db: db,
	}
}

func (p *partnerRepository) GetAllPartners(ctx context.Context) ([]model.Partner, error) {
	query := "select url, description from partners"
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("error getting partners: %v", err)
		return nil, err
	}
	defer rows.Close()

	partners := make([]model.Partner, 0)
	for rows.Next() {
		var partner model.Partner
		if err := rows.Scan(&partner.URL, &partner.Description); err != nil {
			return nil, err
		}
		partners = append(partners, partner)
	}

	return partners, rows.Err()
}
