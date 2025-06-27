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

type SaldoTransferencia struct {
	SaldoOrigem  float32
	SaldoDestino float32
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

func (client FirestoreClient) checkExists(ctx context.Context, documento *string, numeroConta *string) error {
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

func (client FirestoreClient) CheckContaExists(ctx context.Context, fieldName string, valueCampo string) (*domain.Conta, error) {
	// Cria a query para buscar uma conta pelo campo e valor informados
	query := client.client.Collection("conta").Where(fieldName, "==", valueCampo).Limit(1)

	// Executa a consulta
	iter := query.Documents(ctx)

	// Obtém o próximo documento da consulta
	doc, err := iter.Next()

	if err != nil {
		if err == iterator.Done {
			// Se não houver mais documentos (sem resultados)
			log.Warn().Str("field", fieldName).Str("value", valueCampo).Msg("Conta não encontrada")
			return nil, fmt.Errorf("conta com %s '%s' não encontrada", fieldName, valueCampo)
		}
		// Caso ocorra um erro diferente
		log.Error().Err(err).Str("field", fieldName).Str("value", valueCampo).Msg("Erro ao consultar Firestore")
		return nil, fmt.Errorf("erro ao consultar Firestore: %v", err)
	}

	// Mapeia o documento encontrado para o tipo Conta
	var conta domain.Conta
	if err := doc.DataTo(&conta); err != nil {
		log.Error().Err(err).Str("field", fieldName).Str("value", valueCampo).Msg("Erro ao converter dados do documento para conta")
		return nil, fmt.Errorf("erro ao processar dados da conta: %v", err)
	}

	// Retorna a conta encontrada
	return &conta, nil
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

func (client FirestoreClient) GetNumeroConta(id string) (*domain.Conta, error) {
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

func (client FirestoreClient) GetByID(id string) (*domain.Conta, string, error) {
	doc, err := client.client.Collection("conta").
		Where("id", "==", id).
		Documents(context.Background()).Next()
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar documento no Firestore")
		return nil, "", err
	}
	var conta domain.Conta
	if err := doc.DataTo(&conta); err != nil {
		return nil, "", err
	}

	return &conta, doc.Ref.ID, nil
}

func (client FirestoreClient) DeletaConta(id string) (*domain.Conta, error) {
	conta, docID, err := client.GetByID(id)
	if err != nil {
		log.Error().Err(err).Msg("Id não encontrado")
		return nil, err
	}

	ctx := context.Background()
	_, err = client.client.Collection("conta").Doc(docID).Delete(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao deletar documento no Firestore")
		return nil, err
	}

	return conta, nil
}
