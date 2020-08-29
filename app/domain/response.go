package domain

// 定数
const (
	SuccessCode = "00" // 処理成功時
	FailedCode  = "80" // 処理失敗時
)

// Response APIの処理結果に対応する汎用的な構造体
type Response struct {
	ResultCode    string      `json:"resultCode"`
	ResultMessage string      `json:"resultMessage"`
	Contents      interface{} `json:"contents"`
}

// NewResponse Responseのコンストラクタ
func NewResponse() *Response {
	return &Response{
		ResultCode:    FailedCode,
		ResultMessage: "",
		Contents:      nil,
	}
}

// Success 処理成功時
func (r *Response) Success(msg string, contents interface{}) *Response {
	return &Response{
		ResultCode:    SuccessCode,
		ResultMessage: msg,
		Contents:      contents,
	}
}

// Failed 処理失敗時
func (r *Response) Failed(msg string, contents interface{}) *Response {
	return &Response{
		ResultCode:    FailedCode,
		ResultMessage: msg,
		Contents:      contents,
	}
}
