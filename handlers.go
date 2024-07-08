package route

import "net/http"

// DefaultHandler 提供一个默认的HTTP响应处理函数。
func DefaultHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the go-fly route! This is the default handler."))
	})
}

// NotFoundHandler 提供一个自定义的404 Not Found响应处理函数。
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置HTTP状态代码为404 Not Found
		http.NotFound(w, r)
	})
}
