package rest

import (
	"banking/internal/domain"
	"encoding/json"
	"net/http"

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

	responseObject := usuarioToResponse(*newUser)

	json.NewEncoder(response).Encode(responseObject)
}

func UsuarioRequestToDomain(request UsuarioRequest) domain.Usuario{
	objeto := usuarioToDomain(request)	
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

	responseObject := usuarioToResponse(*user)

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
	// lê o id da URL
	vars := mux.Vars(request)
	id := vars["id"]

	// lê o corpo da requisição
	var userRequest UsuarioRequest
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&userRequest)

	// buscar dados atuais no banco
	databaseUser, err := dbClient.GetUsuario(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "008",
			Message: "Erro ao consultar usuário",
		})
		return
	}
	if databaseUser == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "009",
			Message: "Usuário não encontrado",
		})
		return
	}

	// alterar os campos no objeto do banco de acordo com o que veio na requisição
	if userRequest.Nome != "" {
		databaseUser.Nome = userRequest.Nome
	}
	if userRequest.Email != "" {
		databaseUser.Email = userRequest.Email
	}
	if userRequest.Telefone != "" {
		databaseUser.Telefone = userRequest.Telefone
	}

	// salvar no banco
	_, err = dbClient.AlterarUsuario(*databaseUser)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao alterar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "010",
			Message: "Erro ao alterar usuário",
		})
	}	

	responseObject := usuarioToResponse(*databaseUser)

	json.NewEncoder(response).Encode(responseObject)
}






	
