package rest

import (
	"banking/internal/domain"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)



func CadastrarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var user UsuarioRequest

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&user)

	erros := validarDadosUsuario(user)
	if len(erros) > 0 {
		response.WriteHeader(http.StatusBadRequest)
		errosConcatenados := ""
		for _, erro := range erros {
			errosConcatenados += erro + ";"
		}
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "001",
			Message: errosConcatenados,
		})
		return
	}

	newUser, err := dbClient.AddUsuario(UsuarioRequestToDomain(user))
	if err != nil {
		log.Error().Err(err).Msg("Erro ao salvar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "002",
			Message: "Erro ao salvar usuário",
		})
		return
	}
	responseObject := UsuarioResponse{
		Id:       newUser.Id,
		Nome:     newUser.Nome,
		CPF:      newUser.CPF,
		Telefone: newUser.Telefone,
		Email:    newUser.Email,
		CreatedAt: newUser.CreatedAt.Format(time.RFC3339),	
	}
	json.NewEncoder(response).Encode(responseObject)
}

func UsuarioRequestToDomain(request UsuarioRequest) domain.Usuario{
	newUUID := uuid.New().String()		
	objeto := domain.Usuario{
		Id:       newUUID,
		Nome:     request.Nome,
		CPF:      request.CPF,
		Telefone: request.Telefone,
		Email:    request.Email,
		CreatedAt: time.Now(),
	}
	return objeto
}

func ConsultarUsuarioID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]

	user, err := dbClient.GetUsuario(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "003",
			Message: "Erro ao consultar usuário",
		})
		return
	}

	if user == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "004",
			Message: "Usuário não encontrado",
		})
		return
	}

	responseObject := UsuarioResponse{
		Id:       user.Id,
		Nome:     user.Nome,
		CPF:      user.CPF,
		Telefone: user.Telefone,
		Email:    user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),	
	}
	json.NewEncoder(response).Encode(responseObject)

}

func DeletarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]

	user, err := dbClient.DeleteUsuario(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao deletar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "005",
			Message: "Erro ao deletar usuário",
		})
		return
	}

	if user == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "006",
			Message: "Usuário não encontrado",
		})
		return
	}	

		json.NewEncoder(response).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Usuário deletado com sucesso",
		})
	}

func ConsultarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	users, err := dbClient.GetUsuarios()
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar usuários")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "007",
			Message: "Erro ao consultar usuários",
		})
		return
	}

	json.NewEncoder(response).Encode(&users)
}

func AlterarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]

	user, err := dbClient.AlterarUsuario(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao deletar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "005",
			Message: "Erro ao deletar usuário",
		})
		return
	}

	if user == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "006",
			Message: "Usuário não encontrado",
		})
		return
	}	



	




}	