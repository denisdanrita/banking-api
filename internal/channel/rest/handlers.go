package rest

import (
	"banking/internal/domain"
	"encoding/json"
	"net/http"
	"net/mail"
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

func validarDadosUsuario(user UsuarioRequest) []string {
	var erros []string

	if user.Nome == "" {
		erros = append(erros, "Campo nome não preenchido")	
	}

	if user.CPF == "" {
		erros = append(erros, "Campo CPF não preenchido")	
	}

	if user.Telefone == "" {
		erros = append(erros, "Campo telefone não preenchido")			
	}

	if user.Email == "" {
		erros = append(erros, "Campo email não preenchido")			
	}

	if len(user.Nome) > 50 {
		erros = append(erros, "Nome deve ter no máximo 50 caracteres")		
	}	

	if len(user.CPF) < 11 {
		erros = append(erros, "CPF deve ter 11 caracteres")		
	}	

	if _, err := mail.ParseAddress(user.Email); err != nil {
		erros = append(erros, "Email inválido")		
	}	
	
	return erros
}

func ConsultarUsuario(response http.ResponseWriter, request *http.Request) {
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