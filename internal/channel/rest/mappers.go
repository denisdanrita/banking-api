package rest

import (
	"banking/internal/domain"
	"time"

	"github.com/google/uuid"
)


func usuarioToResponse(newUser domain.Usuario) UsuarioResponse {
	return UsuarioResponse{ 
	Id:       newUser.Id,
	Nome:     newUser.Nome,
	CPF:      newUser.CPF,
	Telefone: newUser.Telefone,
	Email:    newUser.Email,
	CreatedAt: newUser.CreatedAt.Format(time.RFC3339),	
}
}

func usuarioToDomain(request UsuarioRequest) domain.Usuario{
	newUUID := uuid.New().String()
  return domain.Usuario{
	Id:       newUUID,
	Nome:     request.Nome,
	CPF:      request.CPF,
	Telefone: request.Telefone,
	Email:    request.Email,
	CreatedAt: time.Now(),
}
}
