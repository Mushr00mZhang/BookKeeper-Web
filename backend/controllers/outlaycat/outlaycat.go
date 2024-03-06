package outlaycat

import (
	outlaycat_dto "bookkeeper-backend/dtos/outlaycat"
	"bookkeeper-backend/models/database"
	"bookkeeper-backend/services/outlaycat"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	base := r.PathPrefix("/outlaycats").Subrouter()
	base.HandleFunc("", List).Methods("GET")
	base.HandleFunc("/paged-list", PagedList).Methods("GET")
	base.HandleFunc("/{Id}", Get).Methods("GET")
	base.HandleFunc("", Create).Methods("POST")
	base.HandleFunc("/{Id}", Update).Methods("PUT")
	base.HandleFunc("/{Id}", Delete).Methods("DELETE")
}
func List(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	dto := outlaycat_dto.ParseListDto(values)
	res := outlaycat.List(database.DB, dto)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func PagedList(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	dto := outlaycat_dto.ParsePagedListDto(values)
	res := outlaycat.PagedList(database.DB, dto)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["Id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
	}
	res := outlaycat.Get(database.DB, id)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var dto outlaycat_dto.CreateDto
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("读取Body失败。%s\n", err)
	}
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		log.Printf("反序列化JSON失败。%s\n", err)
	}
	res := outlaycat.Create(database.DB, dto)
	bytes, err = json.Marshal(res)
	if err != nil {
		log.Printf("序列化JSON失败。%s\n", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
func Update(w http.ResponseWriter, r *http.Request) {
	var dto outlaycat_dto.UpdateDto
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("读取Body失败。%s\n", err)
	}
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		log.Printf("反序列化JSON失败。%s\n", err)
	}
	res := outlaycat.Update(database.DB, dto)
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
	id, err := uuid.Parse(vars["Id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
	}
	res := outlaycat.Delete(database.DB, id)
	bytes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(200)
}
