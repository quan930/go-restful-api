package entity

////匿名成员
//Body struct{
//	// Required: true
//	Code   int      `json:"code"`
//	// Required: true
//	Msg    string      `json:"msg"`
//	// An optional field name to which this validation applies
//	Data   interface{} `json:"data"`
//}

// swagger:response response
type Response struct {
	// response
	// in: body
	Body struct{
		// Required: true
		Code   int      `json:"code"`
		// Required: true
		Msg    string      `json:"msg"`
		// An optional field name to which this validation applies
		Data   interface{} `json:"data"`
	}
}

func NewResponse(Code int,Msg  string,Data interface{}) *Response {
	responseBody := new(Response)
	responseBody.Body.Code = Code
	responseBody.Body.Msg = Msg
	responseBody.Body.Data = Data
	return responseBody
}