package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{} // string -> any

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
	}
}

// PostForm 提取post请求中的参数（位于一个JSON中）
func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

// Query 提取get请求中的参数（位于请求路径）
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// Status 设置Response中的状态码
func (c *Context) Status(status int) {
	c.StatusCode = status
	c.Writer.WriteHeader(status)
}

// SetHeader 设置返回头中的某个字段
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 以纯文本形式返回响应
func (c *Context) String(status int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(status)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 以JSON形式返回响应
func (c *Context) JSON(status int, v interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(status)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(v); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 以二进制数据形式返回响应
func (c *Context) Data(status int, v []byte) {
	c.SetHeader("Content-Type", "application/octet-stream")
	c.Status(status)
	c.Writer.Write(v)
}

// HTML 以HTML形式返回响应
func (c *Context) HTML(status int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(status)
	c.Writer.Write([]byte(html))
}
