package firestore

import (
	"banking/internal/domain"
	"context"

	"cloud.google.com/go/firestore"
	"github.com/rs/zerolog/log"
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





