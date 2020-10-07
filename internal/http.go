package internal 

import (
	"fmt"

	"gopkg.in/resty.v1"
)

var client = resty.New()

// Params 参数类
type Params struct {
	params map[string]string
}

// NewParams 构造
func NewParams() *Params{
	p := &Params{}
	p.params = make(map[string]string)
	return p
}

// Add param to params
func (p *Params) Add(key string, val string) *Params{ 
	p.params[key] = val
	return p
}

// Data params data
func (p *Params) Data() map[string]string{
	return p.params
}

// Get 简单get请求
func Get(url string, params Params) {
	resp, err := client.R(). 
		SetQueryParams(params.params).
		Get(url)
	fmt.Println(resp)
	fmt.Println(err)
}

