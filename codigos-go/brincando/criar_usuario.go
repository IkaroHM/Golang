package Criar_usuario

import (
	"fmt"

	Criar_usuario "github.com/IkaroHM/codigos-go/brincando"
	"github.com/google/uuid"
)


type usuario struct {
	nome string
	senha string
	email string
	uuid string
}

func (f usuario) Criar_usuario_usuario (nome, senha, email, uuid string) {
	f.nome = nome
	f.senha = senha
	f.email = email
	f.uuid = uuid
}

func main() {
	id := uuid.New().String()
	Criar_usuario("Ikaro", )
}