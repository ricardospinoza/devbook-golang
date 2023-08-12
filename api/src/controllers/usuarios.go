package controllers

import "net/http"

// CriarUsuario insere no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário!"))
}

// BuscarUsuarios recupera o usuario no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bucandos todos os Usuário!"))
}

// BuscarUsuario recupera o usuário especifico no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscar  Usuário!"))
}

// AtualizarUsuario atualiza o usuário no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar Usuário!"))
}

// DeletarUsuario remove o usuário no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deeltar Usuário!"))
}
