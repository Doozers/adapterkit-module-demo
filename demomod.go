package demomod

import "context"

type service struct {
	UnimplementedDemomodSvcServer // required by protoc.

	param string
}

func New(param string) DemomodSvcServer {
	return &service{param: param}
}

func (svc service) Sum(ctx context.Context, req *Sum_Request) (*Sum_Result, error) {
	a := req.GetA()
	b := req.GetB()
	sum := a + b
	ret := Sum_Result{C: sum}
	return &ret, nil
}

func (svc *service) EchoStream(stream DemomodSvc_EchoStreamServer) error {
	panic("not implemented")
}

func (svc *service) SayHello(ctx context.Context, _ *Empty) (*HelloResult, error) {
	ret := HelloResult{Msg: "hello! " + svc.param}
	return &ret, nil
}
