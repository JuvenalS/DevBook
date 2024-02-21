package controllers

import "net/http"

// Insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

// Busca todos os usuários cadastrados no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

// Busca um usuário cadastrado no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

// Altera dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// Deleta um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
