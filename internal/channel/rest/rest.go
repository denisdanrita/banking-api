package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func NewServer() {
	router := mux.NewRouter()

	router.HandleFunc("/usuario", CadastrarUsuario).Methods(http.MethodPost)

	log.Info().Msg("Starting server on port 8100")

	err := http.ListenAndServe(":8100", router)
	if err != nil {
		println(err)
	}
}
