package usecase

import (
	"bas-backend/domain/dataproviders"
	"bas-backend/domain/model"
	"context"
)

type PartnerUsecase interface {
	Fetch(ctx context.Context) ([]model.Partner, error)
}

type partnerUsecase struct {
	partnerRepo model.PartnerRepository
}

func NewPartnerUsecase(providers dataproviders.Providers) PartnerUsecase {

	return &partnerUsecase{
		partnerRepo: providers.GetPartnerRepository(),
	}
}

func (pu *partnerUsecase) Fetch(ctx context.Context) ([]model.Partner, error) {
	return pu.partnerRepo.GetAllPartners(ctx)
}
