package router

import "github.com/gorilla/mux"

//Gerar vai retornar as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
