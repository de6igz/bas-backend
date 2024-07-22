package usecase

import (
	"bas-backend/domain/dataproviders"
	"bas-backend/domain/dataproviders/repository"
	"bas-backend/domain/model"
	"context"
)

type DocumentsUsecase interface {
	GetAllDocuments(ctx context.Context) ([]model.Document, error)
}

type docUsecase struct {
	documentRepo repository.DocRepository
}

func NewDocumentUsecase(providers dataproviders.Providers) DocumentsUsecase {

	return &docUsecase{
		documentRepo: providers.GetDocumentRepository(),
	}
}

func (u *docUsecase) GetAllDocuments(ctx context.Context) ([]model.Document, error) {
	docs, err := u.documentRepo.GetAllDocs(ctx)
	if err != nil {
		return nil, err
	}

	return docs, nil
}
