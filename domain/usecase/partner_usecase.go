package usecase

import (
	"bas-backend/domain/model"
	"context"
)

type PartnerUsecase interface {
	Fetch(ctx context.Context) ([]model.Partner, error)
}

type partnerUsecase struct {
	partnerRepo model.PartnerRepository
}

func NewPartnerUsecase(pr model.PartnerRepository) PartnerUsecase {
	return &partnerUsecase{
		partnerRepo: pr,
	}
}

func (pu *partnerUsecase) Fetch(ctx context.Context) ([]model.Partner, error) {
	return pu.partnerRepo.GetAll(ctx)
}
