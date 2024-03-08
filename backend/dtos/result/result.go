package result

import (
	"encoding/json"
	"net/http"
)

type Dto[T any] struct {
	Code   int8   // 状态码
	Result T      // 结果
	Tip    string // 提示
	Error  error  // 错误
}

// 返回结果
func (dto Dto[T]) Return(w http.ResponseWriter, statusCode int) {
	bytes := dto.Marshal()
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
	w.WriteHeader(statusCode)
}

// 序列化
func (dto Dto[T]) Marshal() []byte {
	bytes, err := json.Marshal(dto)
	if err != nil {
		dto = Dto[T]{
			Code:  -1,
			Tip:   "序列化JSON失败",
			Error: err,
		}
		bytes, _ = json.Marshal(dto)
	}
	return bytes
}
