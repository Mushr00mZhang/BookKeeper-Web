package outlay

import (
	outlay_dto "bookkeeper-backend/dtos/outlay"
	"bookkeeper-backend/dtos/pagedlist"
	"bookkeeper-backend/dtos/result"
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
	list, code, err := outlay.List(database.DB, dto)
	res := result.Dto[[]outlay_dto.Dto]{
		Code:   code,
		Result: list,
		Tip:    "查询成功",
		Error:  err,
	}
	if err == nil {
		res.Tip = "查询失败"
		res.Return(w, 500)
		return
	}
	res.Return(w, 200)
}
func PagedList(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	dto := outlay_dto.ParsePagedListDto(values)
	list, code, err := outlay.PagedList(database.DB, dto)
	res := result.Dto[pagedlist.Dto[outlay_dto.Dto]]{
		Code:   code,
		Result: list,
		Tip:    "查询成功",
		Error:  err,
	}
	if err == nil {
		res.Tip = "查询失败"
		res.Return(w, 500)
		return
	}
	res.Return(w, 200)
}
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
		res := result.Dto[*outlay_dto.Dto]{
			Code:  1,
			Tip:   "解析Id失败",
			Error: err,
		}
		res.Return(w, 400)
		return
	}
	item, code, err := outlay.Get(database.DB, id)
	res := result.Dto[*outlay_dto.Dto]{
		Code:   0,
		Result: item,
		Tip:    "查询成功",
		Error:  err,
	}
	if err != nil {
		res.Code = code + 1
		res.Tip = "查询失败"
		res.Return(w, 500)
		return
	}
	res.Return(w, 200)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var dto outlay_dto.CreateDto
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("读取Body失败。%s\n", err)
		res := result.Dto[*uuid.UUID]{
			Code:  1,
			Tip:   "读取Body失败",
			Error: err,
		}
		res.Return(w, 500)
		return
	}
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		log.Printf("反序列化JSON失败。%s\n", err)
		res := result.Dto[*uuid.UUID]{
			Code:  2,
			Tip:   "反序列化JSON失败",
			Error: err,
		}
		res.Return(w, 400)
		return
	}
	id, code, err := outlay.Create(database.DB, dto)
	res := result.Dto[*uuid.UUID]{
		Code:   0,
		Result: id,
		Tip:    "查询成功",
		Error:  err,
	}
	if err != nil {
		res.Code = code + 2
		res.Tip = "查询失败"
		res.Return(w, 500)
		return
	}
	res.Return(w, 200)
}
func Update(w http.ResponseWriter, r *http.Request) {
	var dto outlay_dto.UpdateDto
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
		res := result.Dto[bool]{
			Code:  1,
			Tip:   "解析Id失败",
			Error: err,
		}
		res.Return(w, 400)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("读取Body失败。%s\n", err)
		res := result.Dto[bool]{
			Code:  2,
			Tip:   "读取Body失败",
			Error: err,
		}
		res.Return(w, 500)
		return
	}
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		log.Printf("反序列化JSON失败。%s\n", err)
		res := result.Dto[bool]{
			Code:  3,
			Tip:   "反序列化JSON失败",
			Error: err,
		}
		res.Return(w, 400)
		return
	}
	if id != dto.Id {
		log.Printf("Query与Body中的Id不一致。\n")
		res := result.Dto[bool]{
			Code:  4,
			Tip:   "Query与Body中的Id不一致",
			Error: err,
		}
		res.Return(w, 400)
		return
	}
	s, code, err := outlay.Update(database.DB, dto)
	res := result.Dto[bool]{
		Code:   0,
		Result: s,
		Tip:    "更新成功",
		Error:  err,
	}
	if err != nil {
		res.Code = code + 4
		res.Tip = "更新失败"
		res.Return(w, 500)
		return
	}
	res.Return(w, 200)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("解析Id失败。\n")
		res := result.Dto[bool]{
			Code:   1,
			Result: false,
			Tip:    "解析Id失败",
			Error:  err,
		}
		res.Return(w, 400)
		return
	}
	s, code, err := outlay.Delete(database.DB, id)
	res := result.Dto[bool]{
		Code:   0,
		Result: s,
		Tip:    "删除成功",
		Error:  err,
	}
	if err != nil {
		res.Code = code + 1
		res.Tip = "删除失败"
		res.Return(w, 500)
		return
	}
	res.Return(w, 200)
}
