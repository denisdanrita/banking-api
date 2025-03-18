package rest

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
	Endereço string `json:"endereco,omitempty"`
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
