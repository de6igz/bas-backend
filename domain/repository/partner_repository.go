package repository

import (
	"bas-backend/domain/model"
)

type partnerRepository struct {
	partners []*model.Partner
}

func NewPartnerRepository() model.PartnerRepository {
	return &partnerRepository{
		partners: []*model.Partner{},
	}
}

func (pr *partnerRepository) GetAll() ([]*model.Partner, error) {
	// Здесь нужно добавить логику для работы с базой данных
	return pr.partners, nil
}
