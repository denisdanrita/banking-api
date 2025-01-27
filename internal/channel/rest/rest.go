package rest

import (
	"banking/internal/integration/firestore"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)
var dbClient *firestore.FirestoreClient

func NewServer() {
	dbClient = firestore.NewConnection()
	router := mux.NewRouter()

	router.HandleFunc("/usuario", CadastrarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuario/{id}", ConsultarUsuario).Methods(http.MethodGet)

	log.Info().Msg("Starting server on port 8100")

	err := http.ListenAndServe(":8100", router)
	if err != nil {
		println(err)
	}
}
