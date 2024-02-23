package ziface

type IServer interface {
	// Start start the server, listen and accept connections.
	Start()

	// Stop stop the server, clean the resources.
	Stop()

	// Serve keep the server continuously running,
	Serve()

	// AddRouter add a router for the server,
	// and processed by connection.
	// AddRouter need to be exposed outside the framework.
	AddRouter(IRouter)
}
