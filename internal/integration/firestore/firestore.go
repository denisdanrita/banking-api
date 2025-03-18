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

func NewConnection() *FirestoreClient{
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

	








