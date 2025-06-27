package service

import (
	"banking/internal/domain"
	"banking/internal/integration/firestore"
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type TransferService interface {
	TransferirConta(data domain.Transferencia) (*firestore.SaldoTransferencia, error)
}

type transferServiceImpl struct {
	dbClient *firestore.FirestoreClient
}

func NewTransferService(dbClient *firestore.FirestoreClient) TransferService {
	return transferServiceImpl{
		dbClient: dbClient,
	}
}

func (t transferServiceImpl) TransferirConta(data domain.Transferencia) (*firestore.SaldoTransferencia, error) {
	ctx := context.Background()
	contaOrigem, err := t.dbClient.CheckContaExists(ctx, "numero_conta", data.NumeroContaOrigem)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar a conta de origem")
		return nil, err
	}

	contaDestino, err := t.dbClient.CheckContaExists(ctx, "numero_conta", data.NumeroContaDestino)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar a conta de destino")
		return nil, err
	}

	if contaOrigem.Saldo < data.ValorTransferencia {
		log.Error().Msg("Saldo insuficiente")
		return nil, fmt.Errorf("saldo insuficiente")
	}

	novoSaldocontaOrigem := contaOrigem.Saldo - data.ValorTransferencia
	novoSaldocontaDestino := contaDestino.Saldo + data.ValorTransferencia

	contaOrigemToUpdate := domain.Conta{
		Id:    contaOrigem.Id,
		Saldo: novoSaldocontaOrigem,
	}

	alterarSaldoContaOrigem, err := t.dbClient.AlterarSaldoConta(contaOrigemToUpdate)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao alterar saldo da conta de origem")
		return nil, err
	}

	log.Info().Msg(fmt.Sprintf("Saldo da conta de origem alterado: %f", alterarSaldoContaOrigem.Saldo))

	contaDestinoToUpdate := domain.Conta{
		Id:    contaDestino.Id,
		Saldo: novoSaldocontaDestino,
	}

	alterarSaldoContaDestino, err := t.dbClient.AlterarSaldoConta(contaDestinoToUpdate)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao alterar saldo da conta de destino")
		return nil, err
	}

	log.Info().Msg(fmt.Sprintf("Saldo da conta de destino alterado: %f", alterarSaldoContaDestino.Saldo))

	return &firestore.SaldoTransferencia{
		SaldoOrigem:  novoSaldocontaOrigem,
		SaldoDestino: novoSaldocontaDestino,
	}, nil
}
