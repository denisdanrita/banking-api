package domain

type Usuario struct {
	Nome     string `json:"nome,omitempty"`
	CPF      string `json:"cpf"`
	Telefone string `json:"telefone"`
	Email    string `json:"email,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
