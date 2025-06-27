package service

import (
	"banking/internal/domain"
	"banking/internal/integration/firestore"
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type SaqueService interface {
	SaqueConta(data domain.Saque) (*float32, error)
}

type saqueServiceImpl struct {
	dbClient *firestore.FirestoreClient
}

func NewSaqueService(dbClient *firestore.FirestoreClient) SaqueService {
	return saqueServiceImpl{
		dbClient: dbClient,
	}
}

func (t saqueServiceImpl) SaqueConta(data domain.Saque) (*float32, error) {
	ctx := context.Background()
	conta, err := t.dbClient.CheckContaExists(ctx, "numero_conta", data.NumeroConta)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar a conta")
		return nil, err
	}

	if conta.Saldo < data.ValorSaque {
		log.Error().Msg("Saldo insuficiente")
		return nil, fmt.Errorf("saldo insuficiente")
	}

	conta.Saldo = conta.Saldo - data.ValorSaque
	contaToUpdate := domain.Conta{
		Id:    conta.Id,
		Saldo: conta.Saldo,
	}

	alterarSaldo, err := t.dbClient.AlterarSaldoConta(contaToUpdate)
	log.Info().Msg(fmt.Sprintf("Saldo alterado: %f", alterarSaldo.Saldo))
	if err != nil {
		log.Error().Err(err).Msg("Erro ao alterar o saldo")
		return nil, err
	}
	return &conta.Saldo, nil
}
