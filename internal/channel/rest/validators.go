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

	if user.Endereco == "" {
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

func validarDadosConta(conta CriacaoContaRequest) []string {
	var erros []string

	if conta.Documento == "" {
		erros = append(erros, "Campo documento não preenchido")
	}

	if conta.Nome == "" {
		erros = append(erros, "Campo nome não preenchido")
	}

	if conta.TipoPessoa == "" {		
		erros = append(erros, "Campo tipo de pessoa não preenchido")
	}

	if conta.CodigoBanco == "" {
		erros = append(erros, "Campo código do banco não preenchido")
	}

	if conta.Agencia == "" {
		erros = append(erros, "Campo agência não preenchido")
	}

	if conta.TipoConta == "" {	
		erros = append(erros, "Campo tipo de conta não preenchido")
	}

	if len(conta.Nome) > 50 {
		erros = append(erros, "Nome deve ter no máximo 50 caracteres")
	}

	if conta.TipoPessoa == "F" && len(conta.Documento) < 11 {
		erros = append(erros, "Documento deve ter 11 caracteres")
	}
	
	if conta.TipoPessoa == "J" &&	len(conta.Documento) < 14 {
		erros = append(erros, "Documento deve ter 14 caracteres")
	}

	if _, err := mail.ParseAddress(conta.EmailTitular); err != nil {
		erros = append(erros, "Email inválido")
	}

	return erros
}

func validarDadosDeposito(deposito DepositoContaRequest) []string {
	var erros []string
	if deposito.NumeroConta == "" {
		erros = append(erros, "Campo número da conta não preenchido")
	}
	if deposito.ValorDeposito == 0 {
		erros = append(erros, "Campo valor do depósito não preenchido")
	}
	if deposito.ValorDeposito == 0 {	
		erros = append(erros, "Campo valor do depósito deve ser maior que 0")
	}
	return erros
}

func validarDadosSaque(saque SaqueContaRequest) []string {
	var erros []string
	if saque.NumeroConta == "" {
		erros = append(erros, "Campo número da conta não preenchido")
	}
	if saque.ValorSaque == 0 {
		erros = append(erros, "Campo valor do depósito não preenchido")
	}
	return erros
}
