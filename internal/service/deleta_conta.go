package service

import (
	"banking/internal/domain"
	"banking/internal/integration/firestore"
)

type DeletaContaService interface {
	DeletaConta(id string) (*domain.Conta, error)
}

type deletaContaServiceImpl struct {
	repository *firestore.FirestoreClient
}

func NewDeletaContaService(dbClient *firestore.FirestoreClient) DeletaContaService {
	return deletaContaServiceImpl{
		repository: dbClient,
	}
}

func (t deletaContaServiceImpl) DeletaConta(id string) (*domain.Conta, error) {
	conta, err := t.repository.DeletaConta(id)
	if err != nil {
		return nil, err
	}
	return conta, nil
}
