package usecase

import "bas-backend/domain/model"

type PartnerUsecase interface {
	Fetch() ([]*model.Partner, error)
}

type partnerUsecase struct {
	partnerRepo model.PartnerRepository
}

func NewPartnerUsecase(pr model.PartnerRepository) PartnerUsecase {
	return &partnerUsecase{
		partnerRepo: pr,
	}
}

func (pu *partnerUsecase) Fetch() ([]*model.Partner, error) {
	return pu.partnerRepo.GetAll()
}
