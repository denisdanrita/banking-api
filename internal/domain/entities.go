package domain

import "time"

type Usuario struct {
	Id 		    string `firestore:"id"`
	Nome      string `firestore:"nome"`
	CPF       string `firestore:"cpf"`
	Telefone  string `firestore:"telefone"`
	Email     string `firestore:"email"` 
	CreatedAt time.Time `firestore:"created_at"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
