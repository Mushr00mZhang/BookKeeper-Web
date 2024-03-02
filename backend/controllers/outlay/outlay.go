package outlay

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	base := r.PathPrefix("/outlays").Subrouter()
	base.HandleFunc("", List).Methods("GET")
	base.HandleFunc("/paged-list", PagedList).Methods("GET")
	base.HandleFunc("/{id}", Get).Methods("GET")
	base.HandleFunc("", Create).Methods("POST")
	base.HandleFunc("/{id}", Update).Methods("PUT")
	base.HandleFunc("/{id}", Delete).Methods("DELETE")
}
func List(w http.ResponseWriter, r *http.Request) {
}
func PagedList(w http.ResponseWriter, r *http.Request) {

}
func Get(w http.ResponseWriter, r *http.Request) {

}
func Create(w http.ResponseWriter, r *http.Request) {

}
func Update(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
