package route

// NewRouter 创建并返回一个新的Router实例。
func NewRouter() *Router {
	return &Router{
		root: NewNode(),
	}
}
