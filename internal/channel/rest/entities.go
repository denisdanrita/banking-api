package rest

type TipoPessoa string
const (
	TipoPessoaFisica TipoPessoa = "FISICA"
	TipoPessoaJuridica TipoPessoa = "JURIDICA"
)

type TipoConta string
const (
	TipoContaCorrente TipoConta = "CORRENTE"
	TipoContaPoupanca TipoConta = "POUPANCA"
)	

type UsuarioRequest struct {
	Nome     string `json:"nome,omitempty"`
	CPF      string `json:"cpf"`
	Telefone string `json:"telefone"`
	Email    string `json:"email,omitempty"`
	UserName string `json:"username,omitempty"`
	Senha    string `json:"senha,omitempty"`
}

type UsuarioResponse struct {
	Id        string `json:"id,omitempty"`
	Nome      string `json:"nome,omitempty"`
	CPF       string `json:"cpf"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Token     string `json:"token,omitempty"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e ResponseError) Error() string {
	return e.Code + ": " + e.Message
}

type ClienteRequest struct {
	Nome     string `json:"nome,omitempty"`
	CPF      string `json:"cpf"`
	Telefone string `json:"telefone"`
	Email    string `json:"email,omitempty"`
	Endereco string `json:"endereco,omitempty"`
}

type ClienteResponse struct {
	Id        string `json:"id,omitempty"`
	Nome      string `json:"nome,omitempty"`
	CPF       string `json:"cpf"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email,omitempty"`
	Endereço  string `json:"endereco,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Token     string `json:"token,omitempty"`
}

type CriacaoContaRequest struct {
	CodigoBanco      string `json:"codigo_banco,omitempty"`
	Agencia          string `json:"agencia,omitempty"`
	DigitoAgencia    string `json:"digito_agencia,omitempty"`
	NumeroConta      string `json:"numero_conta,omitempty"`
	DigitoConta      string `json:"digito_conta,omitempty"`
	TipoConta        TipoConta `json:"tipo_conta,omitempty"`
	TipoPessoa			 TipoPessoa `json:"tipo_pessoa,omitempty"`
	Nome             string `json:"nome,omitempty"`
	Documento        string `json:"documento,omitempty"`
	EmailTitular     string `json:"email_titular,omitempty"`
	TelefoneTitular  string `json:"telefone_titular,omitempty"`
	Saldo            float32 `json:"saldo,omitempty"`
}

type CriacaoContaResponse struct {
	Id              string `json:"id,omitempty"`
	CodigoBanco     string `json:"codigo_banco,omitempty"`
	Agencia         string `json:"agencia,omitempty"`
	DigitoAgencia   string `json:"digito_agencia,omitempty"`
	NumeroConta     string `json:"numero_conta,omitempty"`
	DigitoConta     string `json:"digito_conta,omitempty"`
	TipoConta       string `json:"tipo_conta,omitempty"`
	TipoPessoa			string `json:"tipo_pessoa,omitempty"`
	Nome            string `json:"nome,omitempty"`
	Documento       string `json:"documento,omitempty"`
	EmailTitular    string `json:"email_titular,omitempty"`
	TelefoneTitular string `json:"telefone_titular,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
}

type AlteracaoContaRequest struct {	
	Agencia          string `json:"agencia,omitempty"`
	DigitoAgencia    string `json:"digito_agencia,omitempty"`
	TipoConta        TipoConta `json:"tipo_conta,omitempty"`
	TipoPessoa			 TipoPessoa `json:"tipo_pessoa,omitempty"`
	Nome             string `json:"nome,omitempty"`	
	EmailTitular     string `json:"email_titular,omitempty"`
	TelefoneTitular  string `json:"telefone_titular,omitempty"`
	Senha					   string `json:"senha,omitempty"`
}

type ConsultaSaldoResponse struct {
	Saldo           float32 `json:"saldo,omitempty"`
}

type DepositoContaRequest struct {
	NumeroConta      string `json:"numero_conta,omitempty"`
	ValorDeposito		 float32 `json:"valor_deposito,omitempty"`
}

type DepositoContaResponse struct {	
	NumeroConta     string `json:"numero_conta,omitempty"`
	ValorDeposito		float32 `json:"valor_deposito,omitempty"`
	SaldoAlterado 	string `json:"saldo_alterado,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
}

type SaqueContaRequest struct {
	NumeroConta      string `json:"numero_conta,omitempty"`
	ValorSaque  		 float32 `json:"valor_saque,omitempty"`
}

type SaqueContaResponse struct {		
	ValorSaque		  float32 `json:"valor_saque,omitempty"`
	SaldoAlterado 	string `json:"saldo_alterado,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
}