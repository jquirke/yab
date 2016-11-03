// Code generated by thriftrw-plugin-yarpc
// @generated

package fooserver

import (
	"github.com/yarpc/yab/testdata/yarpc/integration/service/foo"
	"golang.org/x/net/context"
	"go.uber.org/thriftrw/protocol"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc"
	"go.uber.org/thriftrw/wire"
)

// Interface is the server-side interface for the Foo service.
type Interface interface {
	Bar(
		ctx context.Context,
		reqMeta yarpc.ReqMeta,
		Arg *int32,
	) (int32, yarpc.ResMeta, error)
}

// New prepares an implementation of the Foo service for
// registration.
//
// 	handler := FooHandler{}
// 	thrift.Register(dispatcher, fooserver.New(handler))
func New(impl Interface) thrift.Service {
	return service{handler{impl}}
}

type service struct{ h handler }

func (service) Name() string {
	return "Foo"
}

func (service) Protocol() protocol.Protocol {
	return protocol.Binary
}

func (s service) Handlers() map[string]thrift.Handler {
	return map[string]thrift.Handler{
		"bar": thrift.HandlerFunc(s.h.Bar),
	}
}

type handler struct{ impl Interface }

func (h handler) Bar(
	ctx context.Context,
	reqMeta yarpc.ReqMeta,
	body wire.Value,
) (thrift.Response, error) {
	var args foo.BarArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, resMeta, err := h.impl.Bar(ctx, reqMeta, args.Arg)

	hadError := err != nil
	result, err := foo.BarHelper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body = result
	}
	return response, err
}
