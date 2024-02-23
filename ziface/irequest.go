package ziface

// IRequest pack the connection and data.
type IRequest interface {
	GetConnection() IConnection

	GetData() []byte
}
