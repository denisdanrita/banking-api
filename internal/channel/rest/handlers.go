package rest

import (
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

	log.Info().Any("user", user).Msg("Requisição cadastrar usuário")

	erros := validarDadosUsuario(user)
	if len(erros) > 0 {
		response.WriteHeader(http.StatusBadRequest)
		errosConcatenados := ""
		for _, erro := range erros {
			errosConcatenados += erro + ";"
		}
		responseError := ResponseError{
			Code:    "001",
			Message: errosConcatenados,
		}
		json.NewEncoder(response).Encode(responseError)
		log.Info().AnErr("erros", responseError).Msg("Erro ao validar dados do usuário")
		return
	}

	newUser, err := dbClient.AddUsuario(usuarioToDomain(user))
	if err != nil {
		log.Error().Err(err).Msg("Erro ao salvar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "002",
			Message: "Erro ao salvar usuário",
		})
		return
	}

	responseObject := usuarioToResponse(*newUser, true)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", responseObject).Msg("Retorno cadastrar usuário")
}

func ConsultarUsuarioID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Consultar usuário por ID")

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

	responseObject := usuarioToResponse(*user, false)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", responseObject).Msg("Retorno consultar usuário por ID")
}

func DeletarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Deletar usuário")

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
	log.Info().Any("user", user).Msg("Retorno deletar usuário")
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
	log.Info().Any("users", users).Msg("Retorno consultar usuários")
}

func AlterarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	// lê o id da URL
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Alterar usuário")

	// lê o corpo da requisição
	var userRequest UsuarioRequest
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&userRequest)
	log.Info().Any("userRequest", userRequest).Msg("Corpo da requisição")

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

	responseObject := usuarioToResponse(*databaseUser, false)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", response).Msg("Retorno alterar usuário")
}

func AlterarSenha(response http.ResponseWriter, request *http.Request) {

}

func CadastrarCliente(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var cliente ClienteRequest

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&cliente)

	log.Info().Any("user", cliente).Msg("Requisição cadastrar cliente")

	erros := validarDadosCliente(cliente)
	if len(erros) > 0 {
		response.WriteHeader(http.StatusBadRequest)
		errosConcatenados := ""
		for _, erro := range erros {
			errosConcatenados += erro + ";"
		}
		responseError := ResponseError{
			Code:    "001",
			Message: errosConcatenados,
		}
		json.NewEncoder(response).Encode(responseError)
		log.Info().AnErr("erros", responseError).Msg("Erro ao validar dados do cliente")
		return
	}

	newCliente, err := dbClient.AddCliente(clienteToDomain(cliente))
	if err != nil {
		log.Error().Err(err).Msg("Erro ao salvar cliente")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "002",
			Message: "Erro ao salvar cliente",
		})
		return
	}

	responseObject := clienteToResponse(*newCliente, true)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", responseObject).Msg("Retorno cadastrar usuário")
}
