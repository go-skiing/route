package route

import "net/http"

// Get 注册一个处理GET请求的路由。
func (r *Router) Get(path string, handler http.Handler) {
	r.Handle("GET", path, handler)
}

// Post 注册一个处理POST请求的路由。
func (r *Router) Post(path string, handler http.Handler) {
	r.Handle("POST", path, handler)
}

// Put 注册一个处理PUT请求的路由。
func (r *Router) Put(path string, handler http.Handler) {
	r.Handle("PUT", path, handler)
}

// Delete 注册一个处理DELETE请求的路由。
func (r *Router) Delete(path string, handler http.Handler) {
	r.Handle("DELETE", path, handler)
}

// Patch 注册一个处理PATCH请求的路由。
func (r *Router) Patch(path string, handler http.Handler) {
	r.Handle("PATCH", path, handler)
}

// Options 注册一个处理OPTIONS请求的路由。
func (r *Router) Options(path string, handler http.Handler) {
	r.Handle("OPTIONS", path, handler)
}

// Head 注册一个处理HEAD请求的路由。
func (r *Router) Head(path string, handler http.Handler) {
	r.Handle("HEAD", path, handler)
}
