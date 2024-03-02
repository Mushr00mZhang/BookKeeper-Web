package outlaycat

import (
	"bookkeeper/modules/book"
	"bookkeeper/modules/database"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	base := r.PathPrefix("/outlaycats").Subrouter()
	base.HandleFunc("", List).Methods("GET")
	base.HandleFunc("/paged-list", PagedList).Methods("GET")
	base.HandleFunc("/{id}", Get).Methods("GET")
	base.HandleFunc("", Create).Methods("POST")
	base.HandleFunc("/{id}", Update).Methods("PUT")
	base.HandleFunc("/{id}", Delete).Methods("DELETE")
}
func List(w http.ResponseWriter, r *http.Request) {
	res := book.ListOutlayCat(database.DB)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func PagedList(w http.ResponseWriter, r *http.Request) {

}
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := book.GetOutlayCat(database.DB, uuid.MustParse(vars["id"]))
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func Create(w http.ResponseWriter, r *http.Request) {
	res := book.CreateOutlayCat(database.DB)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)

}
func Update(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
