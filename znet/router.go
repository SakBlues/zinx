package znet

import "github.com/SakBlues/zinx/ziface"

var _ ziface.IRouter = (*BaseRouter)(nil)

// BaseRouter is the Base class.
// Custom routers can override this.
type BaseRouter struct {
}

// The function of BaseRouter is because
// some routers don't need PreHandler, PostHandle.
// Custome routers inherit BaseRouter so
// don't need to implement the two methods above.
func (r *BaseRouter) PreHandle(request ziface.IRequest)  {}
func (r *BaseRouter) Handle(request ziface.IRequest)     {}
func (r *BaseRouter) PostHandle(request ziface.IRequest) {}
