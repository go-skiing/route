package route

import (
	"context"
	"net/http"
	"strings"
	"sync"
)

// Router 结构体表示一个HTTP路由器，负责管理所有路由并分发请求。
type Router struct {
	root        *Node
	lock        sync.Mutex
	middlewares []Middleware
}

// Node 表示路由树中的一个节点。
type Node struct {
	children map[string]*Node
	handlers map[string]http.Handler
}

// NewNode 创建并返回一个新的节点实例。
func NewNode() *Node {
	return &Node{children: make(map[string]*Node)}
}

// Handle 注册一个新的路由处理函数到指定的HTTP方法和路径。
func (r *Router) Handle(method, path string, handler http.Handler) {
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, mw := range r.middlewares {
		handler = mw(handler)
	}
	parts := strings.Split(path, "/")
	r.root.insert(method, parts, handler)
}

// Use 添加全局中间件到路由器。
func (r *Router) Use(middleware Middleware) {
	r.lock.Lock()
	defer r.lock.Unlock()
	// 添加中间件到中间件列表
	r.middlewares = append(r.middlewares, middleware)
}

// ServeHTTP 实现http.Handler接口，用于接收和处理所有HTTP请求。
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, params := r.root.search(req.Method, req.URL.Path)
	if handler != nil {
		ctx := context.WithValue(req.Context(), "params", params)
		req = req.WithContext(ctx)
		handler.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}

// insert 方法递归地将一个路由的各个部分插入到前缀树中。
func (n *Node) insert(method string, parts []string, handler http.Handler) {
	if len(parts) == 0 {
		n.handlers[method] = handler
		return
	}
	part := parts[0]
	if _, exists := n.children[part]; !exists {
		n.children[part] = NewNode()
	}
	n.children[part].insert(method, parts[1:], handler)
}

// search 在前缀树中查找与给定方法和路径匹配的处理器。
func (n *Node) search(method, path string) (http.Handler, map[string]string) {
	parts := strings.Split(path, "/")
	return n.searchParts(method, parts)
}

// searchParts 递归地搜索处理器。
func (n *Node) searchParts(method string, parts []string) (http.Handler, map[string]string) {
	if len(parts) == 0 {
		handler, exists := n.handlers[method]
		if exists {
			return handler, nil
		}
		return nil, nil
	}
	part := parts[0]
	child, exists := n.children[part]
	if exists {
		return child.searchParts(method, parts[1:])
	}
	return nil, nil
}
