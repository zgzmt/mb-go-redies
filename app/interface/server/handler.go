package server

import (
	"context"
	"net"
)

//tcp连接处理接口，通过实现改接口处理业务

type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Closer() error
}
