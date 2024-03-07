package result

import "encoding/json"

type Dto[T any] struct {
	Status int    // 状态
	Result T      // 结果
	Tip    string // 提示
	Error  error  // 错误
}

// 序列化
func (dto Dto[T]) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(dto)
	if err != nil {
		dto = Dto[T]{
			Status: 0,
			Tip:    "序列化JSON错误。",
			Error:  err,
		}
		bytes, _ = json.Marshal(dto)
	}
	return bytes, err
}
