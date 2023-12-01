package proto

import (
	"github.com/ngrox-io/ngrox/internal/conn"
)

type Protocol interface {
	GetName() string
	WrapConn(conn.Conn, interface{}) conn.Conn
}
