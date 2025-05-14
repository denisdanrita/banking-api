package rest

import (
	"banking/internal/domain"
	"banking/internal/utils"
	"fmt"
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

func contaToDomain(request CriacaoContaRequest) domain.Conta {
	newUUID := uuid.New().String()
	return domain.Conta{
		Id:              newUUID,
		CodigoBanco:     request.CodigoBanco,
		Agencia:         request.Agencia,
		DigitoAgencia:   request.DigitoAgencia,
		NumeroConta:     request.NumeroConta,
		DigitoConta:     request.DigitoConta,
		TipoConta:       string(request.TipoConta),
		TipoPessoa:      string(request.TipoPessoa),
		Nome:            request.Nome,
		Documento:       request.Documento,
		EmailTitular:    request.EmailTitular,
		TelefoneTitular: request.TelefoneTitular,
		Saldo:           request.Saldo,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func contaToResponse(newConta domain.Conta) CriacaoContaResponse {
	contaResponse := CriacaoContaResponse{
		Id:              newConta.Id,
		CodigoBanco:     newConta.CodigoBanco,
		Agencia:         newConta.Agencia,
		DigitoAgencia:   newConta.DigitoAgencia,
		NumeroConta:     newConta.NumeroConta,
		DigitoConta:     newConta.DigitoConta,
		TipoConta:       newConta.TipoConta,
		TipoPessoa:      newConta.TipoPessoa,
		Nome:            newConta.Nome,
		Documento:       newConta.Documento,
		EmailTitular:    newConta.EmailTitular,
		TelefoneTitular: newConta.TelefoneTitular,
		CreatedAt:       newConta.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       newConta.UpdatedAt.Format(time.RFC3339),
	}
	return contaResponse
}

func saldoToResponse(newConta domain.Conta) ConsultaSaldoResponse {
	saldoResponse := ConsultaSaldoResponse{
		Saldo: newConta.Saldo,
	}
	return saldoResponse
}

func depositoToDomain(request DepositoContaRequest) domain.Deposito {
	return domain.Deposito{
		NumeroConta:   request.NumeroConta,
		ValorDeposito: request.ValorDeposito,
	}
}

func depositoToResponse(saldo float32) DepositoContaResponse {
	DepositoResponse := DepositoContaResponse{
		SaldoAlterado: fmt.Sprintf("%.2f", saldo),
	}
	return DepositoResponse
}

func saqueToDomain(request SaqueContaRequest) domain.Saque {
	return domain.Saque{
		NumeroConta: request.NumeroConta,
		ValorSaque:  request.ValorSaque,	
	}
}

func saqueToResponse(saldo float32) SaqueContaResponse {
	SaqueResponse := SaqueContaResponse{
		SaldoAlterado: fmt.Sprintf("%.2f", saldo),	
	}
	return SaqueResponse
}
