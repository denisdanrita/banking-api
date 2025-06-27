package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func CadastrarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var user UsuarioRequest

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&user)

	log.Info().Any("user", user).Msg("Requisição cadastrar usuário")

	erros := validarDadosUsuario(user)
	if len(erros) > 0 {
		response.WriteHeader(http.StatusBadRequest)
		errosConcatenados := ""
		for _, erro := range erros {
			errosConcatenados += erro + ";"
		}
		responseError := ResponseError{
			Code:    "001",
			Message: errosConcatenados,
		}
		json.NewEncoder(response).Encode(responseError)
		log.Info().AnErr("erros", responseError).Msg("Erro ao validar dados do usuário")
		return
	}

	newUser, err := dbClient.AddUsuario(usuarioToDomain(user))
	if err != nil {
		log.Error().Err(err).Msg("Erro ao salvar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "002",
			Message: "Erro ao salvar usuário",
		})
		return
	}

	responseObject := usuarioToResponse(*newUser, true)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", responseObject).Msg("Retorno cadastrar usuário")
}

func ConsultarUsuarioID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Consultar usuário por ID")

	user, err := dbClient.GetUsuario(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "003",
			Message: "Erro ao consultar usuário",
		})
		return
	}

	if user == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "004",
			Message: "Usuário não encontrado",
		})
		return
	}

	responseObject := usuarioToResponse(*user, false)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", responseObject).Msg("Retorno consultar usuário por ID")
}

func DeletarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Deletar usuário")

	user, err := dbClient.DeleteUsuario(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao deletar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "005",
			Message: "Erro ao deletar usuário",
		})
		return
	}

	if user == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "006",
			Message: "Usuário não encontrado",
		})
		return
	}

	json.NewEncoder(response).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Usuário deletado com sucesso",
	})
	log.Info().Any("user", user).Msg("Retorno deletar usuário")
}

func ConsultarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	users, err := dbClient.GetUsuarios()
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar usuários")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "007",
			Message: "Erro ao consultar usuários",
		})
		return
	}

	json.NewEncoder(response).Encode(&users)
	log.Info().Any("users", users).Msg("Retorno consultar usuários")
}

func AlterarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	// lê o id da URL
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Alterar usuário")

	// lê o corpo da requisição
	var userRequest UsuarioRequest
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&userRequest)
	log.Info().Any("userRequest", userRequest).Msg("Corpo da requisição")

	// buscar dados atuais no banco
	databaseUser, err := dbClient.GetUsuario(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "008",
			Message: "Erro ao consultar usuário",
		})
		return
	}
	if databaseUser == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "009",
			Message: "Usuário não encontrado",
		})
		return
	}

	// alterar os campos no objeto do banco de acordo com o que veio na requisição
	if userRequest.Nome != "" {
		databaseUser.Nome = userRequest.Nome
	}
	if userRequest.Email != "" {
		databaseUser.Email = userRequest.Email
	}
	if userRequest.Telefone != "" {
		databaseUser.Telefone = userRequest.Telefone
	}

	// salvar no banco
	_, err = dbClient.AlterarUsuario(*databaseUser)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao alterar usuário")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "010",
			Message: "Erro ao alterar usuário",
		})
	}

	responseObject := usuarioToResponse(*databaseUser, false)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", response).Msg("Retorno alterar usuário")
}

func AlterarSenha(response http.ResponseWriter, request *http.Request) {

}

func CadastrarCliente(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var cliente ClienteRequest

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&cliente)

	log.Info().Any("user", cliente).Msg("Requisição cadastrar cliente")

	erros := validarDadosCliente(cliente)
	if len(erros) > 0 {
		response.WriteHeader(http.StatusBadRequest)
		errosConcatenados := ""
		for _, erro := range erros {
			errosConcatenados += erro + ";"
		}
		responseError := ResponseError{
			Code:    "001",
			Message: errosConcatenados,
		}
		json.NewEncoder(response).Encode(responseError)
		log.Info().AnErr("erros", responseError).Msg("Erro ao validar dados do cliente")
		return
	}

	newCliente, err := dbClient.AddCliente(clienteToDomain(cliente))
	if err != nil {
		log.Error().Err(err).Msg("Erro ao salvar cliente")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "002",
			Message: "Erro ao salvar cliente",
		})
		return
	}

	responseObject := clienteToResponse(*newCliente)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", responseObject).Msg("Retorno cadastrar usuário")
}

func ConsultarClienteID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Consultar cliente por ID")

	cliente, err := dbClient.GetCliente(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar cliente")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "003",
			Message: "Erro ao consultar cliente",
		})
		return
	}

	if cliente == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "004",
			Message: "Cliente não encontrado",
		})
		return
	}

	responseObject := clienteToResponse(*cliente)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("cliente", responseObject).Msg("Retorno consultar cliente por ID")
}

func AlterarCliente(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	// lê o id da URL
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Alterar cliente")

	// lê o corpo da requisição
	var cliRequest ClienteRequest
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&cliRequest)
	log.Info().Any("clienteRequest", cliRequest).Msg("Corpo da requisição")

	if cliRequest.Nome != "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "005",
			Message: "Nome não pode ser alterado",
		})
		return
	}

	if cliRequest.CPF != "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "006",
			Message: "CPF não pode ser alterado",
		})
		return
	}

	// buscar dados atuais no banco
	databaseCliente, err := dbClient.GetCliente(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar cliente")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "007",
			Message: "Erro ao consultar cliente",
		})
		return
	}
	if databaseCliente == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "008",
			Message: "Cliente não encontrado",
		})
		return
	}

	// alterar os campos no objeto do banco de acordo com o que veio na requisição
	if cliRequest.Email != "" {
		databaseCliente.Email = cliRequest.Email
	}
	if cliRequest.Telefone != "" {
		databaseCliente.Telefone = cliRequest.Telefone
	}
	if cliRequest.Endereco != "" {
		databaseCliente.Endereco = cliRequest.Endereco
	}

	// salvar no banco
	_, err = dbClient.AlterarCliente(*databaseCliente)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao alterar cliente")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "009",
			Message: "Erro ao alterar cliente",
		})
	}

	responseObject := clienteToResponse(*databaseCliente)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", response).Msg("Retorno alterar cliente")
}

func DeletarCliente(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Deletar cliente")

	cliente, err := dbClient.DeleteCliente(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao deletar cliente")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "010",
			Message: "Erro ao deletar cliente",
		})
		return
	}

	if cliente == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "011",
			Message: "Cliente não encontrado",
		})
		return
	}

	json.NewEncoder(response).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Cliente deletado com sucesso",
	})
	log.Info().Any("user", cliente).Msg("Retorno deletar cliente")
}

func CadastrarConta(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var conta CriacaoContaRequest

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&conta)

	log.Info().Any("user", conta).Msg("Requisição cadastrar conta")

	erros := validarDadosConta(conta)
	if len(erros) > 0 {
		response.WriteHeader(http.StatusBadRequest)
		errosConcatenados := ""
		for _, erro := range erros {
			errosConcatenados += erro + ";"
		}
		responseError := ResponseError{
			Code:    "001",
			Message: errosConcatenados,
		}
		json.NewEncoder(response).Encode(responseError)
		log.Info().AnErr("erros", responseError).Msg("Erro ao validar dados da conta")
		return
	}

	newConta, err := dbClient.AddConta(contaToDomain(conta))
	if err != nil {
		log.Error().Err(err).Msg("Erro ao salvar conta")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "002",
			Message: "Erro ao salvar conta",
		})
		return
	}

	responseObject := contaToResponse(*newConta)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("conta", responseObject).Msg("Retorno cadastrar conta")
}

func ConsultarConta(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Consultar conta por ID")

	conta, err := dbClient.GetConta(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar conta")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "003",
			Message: "Erro ao consultar conta",
		})
		return
	}

	if conta == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "004",
			Message: "Conta não encontrado",
		})
		return
	}

	responseObject := contaToResponse(*conta)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("conta", responseObject).Msg("Retorno consultar conta por ID")
}

func AlterarConta(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	// lê o id da URL
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Alterar conta")

	// lê o corpo da requisição
	var ctaRequest AlteracaoContaRequest
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&ctaRequest)
	log.Info().Any("clienteRequest", ctaRequest).Msg("Corpo da requisição")

	// buscar dados atuais no banco
	databaseConta, err := dbClient.GetConta(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar conta")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "011",
			Message: "Erro ao consultar conta",
		})
		return
	}
	if databaseConta == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "012",
			Message: "Conta não encontrada",
		})
		return
	}

	// alterar os campos no objeto do banco de acordo com o que veio na requisição
	if ctaRequest.Agencia != "" {
		databaseConta.Agencia = ctaRequest.Agencia
	}
	if ctaRequest.DigitoAgencia != "" {
		databaseConta.DigitoAgencia = ctaRequest.DigitoAgencia
	}
	if ctaRequest.TipoConta != "" {
		databaseConta.TipoConta = string(ctaRequest.TipoConta)
	}
	if ctaRequest.TipoPessoa != "" {
		databaseConta.TipoPessoa = string(ctaRequest.TipoPessoa)
	}
	if ctaRequest.Nome != "" {
		databaseConta.Nome = ctaRequest.Nome
	}
	if ctaRequest.EmailTitular != "" {
		databaseConta.EmailTitular = ctaRequest.EmailTitular
	}
	if ctaRequest.TelefoneTitular != "" {
		databaseConta.TelefoneTitular = ctaRequest.TelefoneTitular
	}

	// salvar no banco
	_, err = dbClient.AlterarConta(*databaseConta)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao alterar conta")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "013",
			Message: "Erro ao alterar conta",
		})
	}

	responseObject := contaToResponse(*databaseConta)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("user", response).Msg("Retorno alterar conta")
}

func DeletarConta(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Deletar conta")

	conta, _, err := dbClient.GetByID(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao deletar conta")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "014",
			Message: "Erro ao deletar conta",
		})
		return
	}

	if conta == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "015",
			Message: "Conta não encontrada",
		})
		return
	}

	json.NewEncoder(response).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Conta deletada com sucesso",
	})
	log.Info().Any("conta", conta).Msg("Retorno deletar conta")
}

func ConsultarSaldo(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	log.Info().Str("id", id).Msg("Consultar saldo da conta")

	conta, err := dbClient.GetConta(id)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao consultar conta")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "003",
			Message: "Erro ao consultar conta",
		})
		return
	}

	if conta == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "004",
			Message: "Conta não encontrada",
		})
		return
	}

	responseObject := saldoToResponse(*conta)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("saldo", responseObject).Msg("Retorno saldo conta")
}

func Depositar(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var deposito DepositoContaRequest

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&deposito)

	log.Info().Any("deposito", deposito).Msg("Requisição deposito conta")

	erros := validarDadosDeposito(deposito)
	if len(erros) > 0 {
		response.WriteHeader(http.StatusBadRequest)
		errosConcatenados := ""
		for _, erro := range erros {
			errosConcatenados += erro + ";"
		}
		responseError := ResponseError{
			Code:    "001",
			Message: errosConcatenados,
		}
		json.NewEncoder(response).Encode(responseError)
		log.Info().AnErr("erros", responseError).Msg("Erro ao validar dados do deposito")
		return
	}

	newDeposito, err := depositoService.DepositoConta(depositoToDomain(deposito))

	if err != nil {
		log.Error().Err(err).Msg("Erro ao realizar deposito")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "002",
			Message: "Erro ao salvar deposito",
		})
		return
	}

	responseObject := depositoToResponse(*newDeposito)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("deposito", responseObject).Msg("Retorno depositar conta")
}

func Sacar(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var saque SaqueContaRequest

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&saque)

	log.Info().Any("saque", saque).Msg("Requisição saque conta")

	erros := validarDadosSaque(saque)
	if len(erros) > 0 {
		response.WriteHeader(http.StatusBadRequest)
		errosConcatenados := ""
		for _, erro := range erros {
			errosConcatenados += erro + ";"
		}
		responseError := ResponseError{
			Code:    "001",
			Message: errosConcatenados,
		}
		json.NewEncoder(response).Encode(responseError)
		log.Info().AnErr("erros", responseError).Msg("Erro ao validar dados do saque")
		return
	}

	newSaque, err := saqueService.SaqueConta(saqueToDomain(saque))

	if err != nil {
		log.Error().Err(err).Msg("Erro ao realizar saque")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "002",
			Message: "Erro ao salvar saque",
		})
		return
	}

	responseObject := saqueToResponse(*newSaque)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("saque", responseObject).Msg("Retorno sacar conta")
}

func Transferir(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var transfer TransferenciaContaRequest

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&transfer)

	log.Info().Any("transfer", transfer).Msg("Requisição transferencia conta")

	saldos, err := transferService.TransferirConta(transferenciaToDomain(transfer))

	if err != nil {
		log.Error().Err(err).Msg("Erro ao realizar transferência")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ResponseError{
			Code:    "002",
			Message: "Erro ao salvar transferência",
		})
		return
	}

	responseObject := transferenciaToResponse(saldos.SaldoDestino, saldos.SaldoOrigem)

	json.NewEncoder(response).Encode(responseObject)
	log.Info().Any("transferir", responseObject).Msg("Retorno transferir conta")
}
