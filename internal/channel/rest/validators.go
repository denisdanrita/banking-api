package rest

import "net/mail"

func validarDadosUsuario(user UsuarioRequest) []string {
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

func validarDadosCliente(user ClienteRequest) []string {
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

	if user.Endereço == "" {
		erros = append(erros, "Campo endereço não preenchido")
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
