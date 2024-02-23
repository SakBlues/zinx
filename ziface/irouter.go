package ziface

// IRouter determine the HandleFunc,
// and then deal with Request with HandleFunc
type IRouter interface {
	PreHandle(IRequest)
	Handle(IRequest)
	PostHandle(IRequest)
}
