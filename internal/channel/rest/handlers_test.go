package rest

import (
	"banking/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestValidarUsuarioOK(t *testing.T) {
	usuario := domain.Usuario{	
		Nome: 	 "João",	
		CPF: 	 "12345678901",
		Telefone: "12345678901",
		Email: "aaaa@gmail.com",
	}
	erros := validarDadosUsuario(usuario)
	assert.Equal(t, 0, len(erros))
}

func TestValidarUsuarioFaltandoCPF(t *testing.T) {
	usuario := domain.Usuario{	
		Nome: 	 "João",		
		Telefone: "12345678901",
		Email: "aaaa@gmail.com",
	}
	erros := validarDadosUsuario(usuario)
	assert.Equal(t, 1, len(erros))
	assert.Equal(t,  "Campo CPF não preenchido", erros[0])
}

func TestValidarUsuarioEmailInvalido(t *testing.T) {
	usuario := domain.Usuario{	
		Nome: 	 "João",	
		CPF: 	 "12345678901",
		Telefone: "12345678901",
		Email: "aaaa",
	}
	erros := validarDadosUsuario(usuario)
	assert.Equal(t, 1, len(erros))
	assert.Equal(t,  "Email inválido", erros[0])
}

func TestValidarUsuarioNomeMaiorQue50(t *testing.T) {
	usuario := domain.Usuario{	
		Nome: 	 "marcosmarcosmarcosmarcosmarcosmarcosmarcosmarcosmarcosmarcosmarcos",
		CPF: 	 "12345678901",
		Telefone: "12345678901",
		Email: "denis@gmail.com",
	}
	erros := validarDadosUsuario(usuario)
	assert.Equal(t, 1, len(erros))
	assert.Equal(t,  "Nome deve ter no máximo 50 caracteres", erros[0])
}
