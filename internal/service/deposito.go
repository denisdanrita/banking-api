package service

import (
	"banking/internal/domain"
	"banking/internal/integration/firestore"
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type DepositoService interface {
	DepositoConta(data domain.Deposito) (*float32, error)
}

type depositoServiceImpl struct {
	dbClient *firestore.FirestoreClient
}

func NewDepositoService(dbClient *firestore.FirestoreClient) DepositoService {
	return depositoServiceImpl{
		dbClient: dbClient,
	}
}

func (t depositoServiceImpl) DepositoConta(data domain.Deposito) (*float32, error) {
	ctx := context.Background()
	conta, err := t.dbClient.CheckContaExists(ctx, "numero_conta", data.NumeroConta)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar a conta")
		return nil, err
	}

	conta.Saldo = conta.Saldo + data.ValorDeposito
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
