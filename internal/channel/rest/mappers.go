package rest

import (
	"banking/internal/domain"
	"banking/internal/utils"
	"time"

	"github.com/google/uuid"
)

var now = time.Now()

func usuarioToResponse(newUser domain.Usuario, returnToken bool) UsuarioResponse {
	usuarioResponse := UsuarioResponse{
		Id:        newUser.Id,
		Nome:      newUser.Nome,
		CPF:       newUser.CPF,
		Telefone:  newUser.Telefone,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt.Format(time.RFC3339),
	}
	if returnToken {
		usuarioResponse.Token = newUser.Token
	}
	return usuarioResponse
}

func usuarioToDomain(request UsuarioRequest) domain.Usuario {
	newUUID := uuid.New().String()
	return domain.Usuario{
		Id:        newUUID,
		Nome:      request.Nome,
		CPF:       request.CPF,
		Telefone:  request.Telefone,
		Email:     request.Email,
		CreatedAt: now,
		UserName:  request.UserName,
		Token:     utils.GenerateToken(request.UserName, request.Senha),
	}
}

func clienteToDomain(request ClienteRequest) domain.Cliente {
	newUUID := uuid.New().String()
	return domain.Cliente{
		Id:        newUUID,
		Nome:      request.Nome,
		CPF:       request.CPF,
		Telefone:  request.Telefone,
		Email:     request.Email,
		Endereco:  request.Endereco,
		CreatedAt: now,
	}
}

func clienteToResponse(newCliente domain.Cliente) ClienteResponse {
	clienteResponse := ClienteResponse{
		Id:        newCliente.Id,
		Nome:      newCliente.Nome,
		CPF:       newCliente.CPF,
		Telefone:  newCliente.Telefone,
		Email:     newCliente.Email,
		Endereço:  newCliente.Endereco,
		CreatedAt: newCliente.CreatedAt.Format(time.RFC3339),
	}
	return clienteResponse
}
