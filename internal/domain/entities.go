package domain

import "time"

type Usuario struct {
	Id        string    `firestore:"id"`
	Nome      string    `firestore:"nome"`
	CPF       string    `firestore:"cpf"`
	Telefone  string    `firestore:"telefone"`
	Email     string    `firestore:"email"`
	UserName  string    `firestore:"username"`
	CreatedAt time.Time `firestore:"created_at"`
	Token     string    `firestore:"token"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Cliente struct {
	Id        string    `firestore:"id"`
	Nome      string    `firestore:"nome"`
	CPF       string    `firestore:"cpf"`
	Telefone  string    `firestore:"telefone"`
	Email     string    `firestore:"email"`
	Endereco  string    `firestore:"endereco"`
	CreatedAt time.Time `firestore:"created_at"`
}

type Conta struct {
	Id              string    `firestore:"id"`
	CodigoBanco     string    `firestore:"codigo_banco"`
	Agencia         string    `firestore:"agencia"`
	DigitoAgencia   string    `firestore:"digito_agencia"`
	NumeroConta     string    `firestore:"numero_conta"`
	DigitoConta     string    `firestore:"digito_conta"`
	TipoConta       string    `firestore:"tipo_conta"`
	TipoPessoa      string    `firestore:"tipo_pessoa"`
	Nome            string    `firestore:"nome"`
	Documento       string    `firestore:"documento"`
	EmailTitular    string    `firestore:"email_titular"`
	TelefoneTitular string    `firestore:"telefone_titular"`
	Saldo           float32   `firestore:"saldo"`
	CreatedAt       time.Time `firestore:"created_at"`
	UpdatedAt       time.Time `firestore:"updated_at"`
}

type Deposito struct {
	Id            string    `firestore:"id"`
	NumeroConta   string    `firestore:"numero_conta"`
	ValorDeposito float32   `firestore:"valor_deposito"`
	SaldoAlterado float32   `firestore:"saldo_alterado"`
	CreatedAt     time.Time `firestore:"created_at"`
}

type Saque struct {
	Id          string    `firestore:"id"`
	NumeroConta string    `firestore:"numero_conta"`
	ValorSaque  float32   `firestore:"valor_saque"`
	CreatedAt   time.Time `firestore:"created_at"`
}
type Transferencia struct {
	Id                 string    `firestore:"id"`
	NumeroContaOrigem  string    `firestore:"numero_conta_origem"`
	NumeroContaDestino string    `firestore:"numero_conta_destino"`
	ValorTransferencia float32   `firestore:"valor_transferencia"`
	CreatedAt          time.Time `firestore:"created_at"`
}
