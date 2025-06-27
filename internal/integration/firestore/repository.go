package firestore

import (
	"banking/internal/domain"
	"context"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (domain.Conta, error)
	DeletaConta(ctx context.Context, id string) (domain.Conta, error)
}
