package rest

type UsuarioRequest struct {
	Nome     string `json:"nome,omitempty"`
	CPF      string `json:"cpf"`
	Telefone string `json:"telefone"`
	Email    string `json:"email,omitempty"`
}

type UsuarioResponse struct {
	Id			 string `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	CPF      string `json:"cpf"`
	Telefone string `json:"telefone"`
	Email    string `json:"email,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}