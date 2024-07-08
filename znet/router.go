package znet

import "zinx/ziface"

type BaseRouter struct {
}

func (router *BaseRouter) PreHandler(request *ziface.IRequest)  {}
func (router *BaseRouter) Handler(request *ziface.IRequest)     {}
func (router *BaseRouter) PostHandler(request *ziface.IRequest) {}
