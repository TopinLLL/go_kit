package endpoint

import (
	"context"
	"errors"
	"go_kit/service"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

//请求模型
type ArithmeticRequest struct {
	Type string `json:"type"`
	A int `json:"a"`
	B int `json:"b"`
}

//应答模型
type ArithmeticResponse struct {
	Result int `json:"result"`
	Error error `json:"error"`
}

//生成EndPoint
func MakeArithmeticEndpoint(srv service.ArithmeticService)endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequest)
		var (
			reqType string
			res,a,b int
			callError error
		)
		a=req.A
		b=req.B
		reqType=req.Type
		switch  {
		case strings.EqualFold(reqType,"add"):
			 res= srv.Add(a, b)
		default:
			return nil,errors.New("method not exist")
		}
		resp:=ArithmeticResponse{
			Result: res,
			Error:  callError,
		}
		return resp,nil
	}
}
