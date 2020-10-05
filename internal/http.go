package handlers

import (
	"fmt"

	"gopkg.in/resty.v1"
)

var client = resty.New()

// Params 参数类
type Params struct {
	params map[string]string
}

// Add param to params
func (p *Params) Add(key string, val string){ 
	p.params[key] = val
}

// Get 简单get请求
func Get(url string, params Params) {
	resp, err := client.R(). 
		SetQueryParams(params.params).
		Get(url)
	fmt.Println(resp)
	fmt.Println(err)
}

