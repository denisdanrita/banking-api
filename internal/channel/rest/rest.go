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
	router.Use(autenticador)

	router.HandleFunc("/usuario", CadastrarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuario", ConsultarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", ConsultarUsuarioID).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", DeletarUsuario).Methods(http.MethodDelete)
	router.HandleFunc("/usuario/{id}", AlterarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuario/{id}/senha", AlterarSenha).Methods(http.MethodPut)
	router.HandleFunc("/cliente", CadastrarCliente).Methods(http.MethodPost)

	log.Info().Msg("Starting server on port 8100")

	err := http.ListenAndServe(":8100", router)
	if err != nil {
		println(err)
	}
}

func autenticador(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		usuario, err := dbClient.GetUsuarioByToken(token)
		if err != nil || usuario == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Info().Str("usuario", usuario.Nome).Msg("Usuário autenticado")

		next.ServeHTTP(w, r)
	})
}
