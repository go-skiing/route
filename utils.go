package route

import (
	"net/http"
	"strings"
)

// Use 添加全局中间件到路由器。
func (r *Route) Use(middleware Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

// insert 方法递归地将一个路由的各个部分插入到前缀树中。
func (n *Node) insert(method string, parts []string, handler http.Handler) {
	if len(parts) == 0 {
		n.handlers[method] = handler
		return
	}
	part := parts[0]
	if part[0] == ':' { // 动态路由参数
		part = ":param"
		n.param = parts[0][1:]
	}
	if _, exists := n.children[part]; !exists {
		n.children[part] = &Node{
			children: make(map[string]*Node),
			handlers: make(map[string]http.Handler),
		}
	}
	n.children[part].insert(method, parts[1:], handler)
}

// search 在前缀树中查找与给定方法和路径匹配的处理器。
func (n *Node) search(method string, parts []string) http.Handler {
	if len(parts) == 0 || (len(parts) == 1 && parts[0] == "") {
		if handler, exists := n.handlers[method]; exists {
			return handler
		}
		return nil
	}
	part := parts[0]
	if child, exists := n.children[part]; exists {
		return child.search(method, parts[1:])
	} else if child, exists := n.children[":param"]; exists {
		// 处理动态路由参数
		return child.search(method, parts[1:])
	}
	return nil
}

// Handle 注册一个新的路由处理函数到指定的HTTP方法和路径。
func (r *Route) Handle(method string, path string, handler http.Handler) {
	for _, mw := range r.middlewares {
		handler = mw(handler)
	}
	parts := strings.Split(path, "/")
	r.root.insert(method, parts, handler)
}

// ServeHTTP 实现http.Handler接口，用于接收和处理所有HTTP请求。
func (r *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler := r.root.search(req.Method, strings.Split(req.URL.Path, "/"))
	if handler != nil {
		handler.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func (r *Route) Get(path string, handler http.Handler)     { r.Handle("GET", path, handler) }
func (r *Route) Post(path string, handler http.Handler)    { r.Handle("POST", path, handler) }
func (r *Route) Put(path string, handler http.Handler)     { r.Handle("PUT", path, handler) }
func (r *Route) Delete(path string, handler http.Handler)  { r.Handle("DELETE", path, handler) }
func (r *Route) Patch(path string, handler http.Handler)   { r.Handle("PATCH", path, handler) }
func (r *Route) Options(path string, handler http.Handler) { r.Handle("OPTIONS", path, handler) }
func (r *Route) Head(path string, handler http.Handler)    { r.Handle("HEAD", path, handler) }
func (r *Route) Connect(path string, handler http.Handler) { r.Handle("CONNECT", path, handler) }
func (r *Route) Trace(path string, handler http.Handler)   { r.Handle("TRACE", path, handler) }
