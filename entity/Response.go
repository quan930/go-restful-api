package entity

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