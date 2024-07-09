package route

import (
	"net/http"
)

// Route 结构体包含路由的根节点和全局中间件。
type Route struct {
	root        *Node
	middlewares []Middleware
}

// Node 结构体代表前缀树中的一个节点。
type Node struct {
	children map[string]*Node
	handlers map[string]http.Handler
	param    string // 动态路由参数名
}

// NewRoute 创建并返回一个新的路由器实例。
func NewRoute() *Route {
	return &Route{
		root: &Node{
			children: make(map[string]*Node),
			handlers: make(map[string]http.Handler),
		},
		middlewares: []Middleware{},
	}
}
