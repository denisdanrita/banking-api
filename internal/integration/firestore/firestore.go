package firestore

import (
	"banking/internal/domain"
	"context"
	"fmt"
	"cloud.google.com/go/firestore"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/iterator"
)

var firestoreClient *FirestoreClient

type FirestoreClient struct {
	client *firestore.Client
}

func NewConnection() *FirestoreClient {
	if firestoreClient != nil {
		return firestoreClient
	}
	ctx := context.Background()
	localClient, err := firestore.NewClientWithDatabase(ctx, "banking-api-444312", "banking")
	if err != nil {
		log.Fatal().Err(err).Msg("Error initializing Firestore")
	}
	firestoreClient = &FirestoreClient{
		client: localClient,
	}
	return firestoreClient
}

func (client FirestoreClient) AddUsuario(data domain.Usuario) (*domain.Usuario, error) {
	_, _, err := client.client.Collection("usuario").Add(context.Background(), data)
	if err != nil {
		log.Error().Err(err).Msg("Error adding document")
		return nil, err
	}
	return &data, nil
}

func (client FirestoreClient) GetUsuario(id string) (*domain.Usuario, error) {
	doc, err := client.client.Collection("usuario").
		Where("id", "==", id).
		Documents(context.Background()).Next()
	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	var user domain.Usuario
	doc.DataTo(&user)
	return &user, nil

}

func (client FirestoreClient) DeleteUsuario(id string) (*domain.Usuario, error) {
	doc, err := client.client.Collection("usuario").
		Where("id", "==", id).
		Documents(context.Background()).Next()
	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	var user domain.Usuario
	doc.DataTo(&user)
	_, err = doc.Ref.Delete(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Error deleting document")
		return nil, err
	}
	return &user, nil

}

func (client FirestoreClient) GetUsuarios() (*[]domain.Usuario, error) {
	iter := client.client.Collection("usuario").Documents(context.Background())
	var users []domain.Usuario
	for {
		doc, err := iter.Next()
		if err != nil {
			log.Error().Err(err).Msg("Error getting")
			break
		}
		var user domain.Usuario
		doc.DataTo(&user)
		users = append(users, user)
	}
	return &users, nil
}

func (client FirestoreClient) AlterarUsuario(data domain.Usuario) (*domain.Usuario, error) {
	doc, err := client.client.Collection("usuario").
		Where("id", "==", data.Id).
		Documents(context.Background()).Next()

	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	_, err = doc.Ref.Update(context.Background(), []firestore.Update{
		{Path: "nome", Value: data.Nome},
		{Path: "email", Value: data.Email},
		{Path: "telefone", Value: data.Telefone},
	})

	if err != nil {
		log.Error().Err(err).Msg("Error updating document")
		return nil, err
	}
	return &data, nil
}

func (client FirestoreClient) GetUsuarioByToken(token string) (*domain.Usuario, error) {
	doc, err := client.client.Collection("usuario").
		Where("token", "==", token).
		Documents(context.Background()).Next()
	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	var user domain.Usuario
	doc.DataTo(&user)
	return &user, nil

}

func (client FirestoreClient) AddCliente(data domain.Cliente) (*domain.Cliente, error) {
	ctx := context.Background()

	if err := client.checkClienteExists(ctx, "cpf", data.CPF); err != nil {
		return nil, err
	}
	_, _, err := client.client.Collection("cliente").Add(context.Background(), &data)
	if err != nil {
		log.Error().Err(err).Msg("Error adding document")
		return nil, err
	}
	return &data, nil
}

func (client FirestoreClient) checkClienteExists(ctx context.Context, fieldName string, value string) error {
	query := client.client.Collection("cliente").Where(fieldName, "==", value).Limit(1).Documents(ctx)
	doc, err := query.Next()
	if err != nil && err != iterator.Done {
		log.Error().Err(err).Str("field", fieldName).Msg("Error querying Firestore for existing client")
		return err
	}

	if doc != nil {
		err := fmt.Errorf("client with %s '%s' already exists", fieldName, value)
		log.Warn().Err(err).Str("field", fieldName).Msg("Client already exists")
		return err
	}

	return nil
}

func (client FirestoreClient) GetCliente(id string) (*domain.Cliente, error) {
	doc, err := client.client.Collection("cliente").
		Where("id", "==", id).
		Documents(context.Background()).Next()
	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	var cliente domain.Cliente
	doc.DataTo(&cliente)
	return &cliente, nil

}

func (client FirestoreClient) AlterarCliente(data domain.Cliente) (*domain.Cliente, error) {
	doc, err := client.client.Collection("cliente").
		Where("id", "==", data.Id).
		Documents(context.Background()).Next()

	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	_, err = doc.Ref.Update(context.Background(), []firestore.Update{
		{Path: "email", Value: data.Email},
		{Path: "telefone", Value: data.Telefone},
		{Path: "endereco", Value: data.Endereco},
	})

	if err != nil {
		log.Error().Err(err).Msg("Error updating document")
		return nil, err
	}
	return &data, nil
}

func (client FirestoreClient) DeleteCliente(id string) (*domain.Cliente, error) {
	doc, err := client.client.Collection("cliente").
		Where("id", "==", id).
		Documents(context.Background()).Next()
	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	var cliente domain.Cliente
	doc.DataTo(&cliente)
	_, err = doc.Ref.Delete(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Error deleting document")
		return nil, err
	}
	return &cliente, nil
}

func (client FirestoreClient) AddConta(data domain.Conta) (*domain.Conta, error) {
	ctx := context.Background()
	err := client.checkExists(ctx, nil, &data.NumeroConta)
	if err != nil {
		log.Error().Err(err).Msg("Conta ja existe")
		return nil, err
	}

	err = client.checkExists(ctx, &data.Documento, nil)
	if err != nil {		
		log.Error().Err(err).Msg("Documento ja existe")
		return nil, err
	}

	_, _, err = client.client.Collection("conta").Add(context.Background(), &data)
	if err != nil {
		log.Error().Err(err).Msg("Error adding document")
		return nil, err
	}
	return &data, nil
}

func (client FirestoreClient) checkExists(ctx context.Context, documento *string, numeroConta *string) (error) {
	if documento != nil {
		return client.checkContaDocExists(ctx, "documento", *documento)
	}
	if numeroConta != nil {
		return client.checkContaDocExists(ctx, "numero_conta", *numeroConta)
	}
	return nil
}


func (client FirestoreClient) checkContaDocExists(ctx context.Context, fieldName string, valueCampo string) error {	
	query := client.client.Collection("conta").Where(fieldName, "==", valueCampo).Limit(1).Documents(ctx)
	doc, err := query.Next()
	if err != nil && err != iterator.Done {
		log.Error().Err(err).Str("field", fieldName).Msg("Error querying Firestore for requested data")
		return err
	}
	if doc != nil {
		err := fmt.Errorf("query with %s '%s' already exists", fieldName, valueCampo)
		log.Warn().Err(err).Str("field", fieldName).Msg("The information consulted already exists")
		return err
	}
	return nil
}

func (client FirestoreClient) GetConta(id string) (*domain.Conta, error) {
	doc, err := client.client.Collection("conta").
		Where("id", "==", id).
		Documents(context.Background()).Next()
	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	var conta domain.Conta
	doc.DataTo(&conta)
	return &conta, nil
}

func (client FirestoreClient) AlterarConta(data domain.Conta) (*domain.Conta, error) {
	doc, err := client.client.Collection("conta").
		Where("id", "==", data.Id).
		Documents(context.Background()).Next()

	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	_, err = doc.Ref.Update(context.Background(), []firestore.Update{
		{Path: "agencia", Value: data.Agencia},
		{Path: "digito_agencia", Value: data.DigitoAgencia},
		{Path: "tipo_conta", Value: data.TipoConta},
		{Path: "tipo_pessoa", Value: data.TipoPessoa},
		{Path: "nome", Value: data.Nome},
		{Path: "email_titular", Value: data.EmailTitular},
		{Path: "telefone_titular", Value: data.TelefoneTitular},
	})

	if err != nil {
		log.Error().Err(err).Msg("Error updating document")
		return nil, err
	}
	return &data, nil
}

func (client FirestoreClient) DeleteConta(id string) (*domain.Conta, error) {
	doc, err := client.client.Collection("conta").
		Where("id", "==", id).
		Documents(context.Background()).Next()
	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	var conta domain.Conta
	doc.DataTo(&conta)
	_, err = doc.Ref.Delete(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Error deleting document")
		return nil, err
	}
	return &conta, nil
}

func (client FirestoreClient) DepositoConta(data domain.Deposito) (*float32, error) {
	ctx := context.Background()
	conta, err := client.checkContaExists(ctx, "numero_conta", data.NumeroConta)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar a conta")
		return nil, err
	}

	conta.Saldo = conta.Saldo + data.ValorDeposito
	contaToUpdate := domain.Conta{
		Id:    conta.Id,
		Saldo: conta.Saldo,
	}
	alterarSaldo, err := client.AlterarSaldoConta(contaToUpdate)
	log.Info().Msg(fmt.Sprintf("Saldo alterado: %f", alterarSaldo.Saldo))
	if err != nil {
		log.Error().Err(err).Msg("Erro ao alterar o saldo")
		return nil, err
	}
	return &conta.Saldo, nil
}

func (client FirestoreClient) checkContaExists(ctx context.Context, fieldName string, valueCampo string) (*domain.Conta, error) {
	query := client.client.Collection("conta").Where(fieldName, "==", valueCampo).Limit(1).Documents(ctx)
	doc, err := query.Next()
	if err != nil && err != iterator.Done {
		log.Error().Err(err).Str("field", "numero_conta").Msg("Error querying Firestore for requested data")
		return nil, err
	}
	if doc != nil {
		var conta domain.Conta
		doc.DataTo(&conta)
		return &conta, nil
	}
	return nil, fmt.Errorf("query with numero_conta '%s' not found", valueCampo)
}

func (client FirestoreClient) AlterarSaldoConta(data domain.Conta) (*domain.Conta, error) {
	doc, err := client.client.Collection("conta").
		Where("id", "==", data.Id).
		Documents(context.Background()).Next()

	if err != nil {
		log.Error().Err(err).Msg("Error getting document")
		return nil, err
	}
	_, err = doc.Ref.Update(context.Background(), []firestore.Update{
			{Path: "saldo", Value: data.Saldo},	
			})	
	if err != nil {
		log.Error().Err(err).Msg("Error updating document")
		return nil, err
	}
	return &data, nil
}

func (client FirestoreClient) SaqueConta(data domain.Saque) (*float32, error) {
	ctx := context.Background()
	conta, err := client.checkContaExists(ctx, "numero_conta", data.NumeroConta)
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

	alterarSaldo, err := client.AlterarSaldoConta(contaToUpdate)
	log.Info().Msg(fmt.Sprintf("Saldo alterado: %f", alterarSaldo.Saldo))
	if err != nil {
		log.Error().Err(err).Msg("Erro ao alterar o saldo")
		return nil, err
	}
	return &conta.Saldo, nil
}