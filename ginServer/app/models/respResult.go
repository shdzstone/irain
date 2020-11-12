package models

//接口响应参数
type RespResult struct {
	Code int           `json:"code"`    // 业务响应状态码
	Msg  string        `json:"message"` // 提示信息
	Data interface{} 	`json:"data"`    // 数据
}
