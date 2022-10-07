package web

type Resp struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) *Resp {
	if code < 300 {
		return &Resp{code, data, ""}
	}
	return &Resp{code, nil, err}
}
