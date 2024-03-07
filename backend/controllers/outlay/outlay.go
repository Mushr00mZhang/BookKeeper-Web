package outlay

import (
	outlay_dto "bookkeeper-backend/dtos/outlay"
	"bookkeeper-backend/models/database"
	"bookkeeper-backend/services/outlay"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
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
	values := r.URL.Query()
	dto := outlay_dto.ParseListDto(values)
	res := outlay.List(database.DB, dto)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func PagedList(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	dto := outlay_dto.ParsePagedListDto(values)
	res := outlay.PagedList(database.DB, dto)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
	}
	res := outlay.Get(database.DB, id)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var dto outlay_dto.CreateDto
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("读取Body失败。%s\n", err)
	}
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		log.Printf("反序列化JSON失败。%s\n", err)
	}
	res := outlay.Create(database.DB, dto)
	bytes, err = json.Marshal(res)
	if err != nil {
		log.Printf("序列化JSON失败。%s\n", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func Update(w http.ResponseWriter, r *http.Request) {
	var dto outlay_dto.UpdateDto
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("读取Body失败。%s\n", err)
	}
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		log.Printf("反序列化JSON失败。%s\n", err)
	}
	res := outlay.Update(database.DB, dto)
	bytes, err = json.Marshal(res)
	if err != nil {
		log.Printf("序列化JSON失败。%s\n", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
	}
	res := outlay.Delete(database.DB, id)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
