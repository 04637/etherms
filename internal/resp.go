package internal

import "log"

// Resp 响应类
type Resp struct {
	code int
	msg  string
	data interface{}
}

// Ok return successful data
func (r *Resp) Ok(data interface{}) Resp {
	return Resp{code: 200, data: data, msg: "ok"}
}

// Fail return failed msg
func (r *Resp) Fail(msg string) Resp {
	return Resp{code: 500, data: nil, msg: msg}
}

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
