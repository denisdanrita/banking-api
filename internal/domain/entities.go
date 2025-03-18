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
	Id        string `firestore:"id"`
	Nome      string `firestore:"nome"`
	CPF       string `firestore:"cpf"`
	Telefone  string `firestore:"telefone"`
	Email     string `firestore:"email"`
	Endereço  string `firestore:"endereco"`
	CreatedAt time.Time `firestore:"created_at"`
}
