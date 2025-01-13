package rest

import (
	"banking/internal/domain"
	"encoding/json"
	"net/http"
	"net/mail"
)

func CadastrarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var user domain.Usuario

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&user)

	erros := validarDadosUsuario(user)
	if len(erros) > 0 {
		response.WriteHeader(http.StatusBadRequest)
		errosConcatenados := ""
		for _, erro := range erros {
			errosConcatenados += erro + ";"
		}
		json.NewEncoder(response).Encode(domain.Error{
			Code:    "001",
			Message: errosConcatenados,
		})
		return
	}
	json.NewEncoder(response).Encode(user)
}

func validarDadosUsuario(user domain.Usuario) []string {
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