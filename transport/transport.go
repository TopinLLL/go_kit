package transport

import (
	"context"
	"encoding/json"
	"go_kit/endpoint"
	"go_kit/service"
	"net/http"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

//Request将JSON反序列化为对象
func DecodeArithmeticRequest(ctx context.Context,req *http.Request)(interface{},error){
	var arithmeticRequest endpoint.ArithmeticRequest
	err := json.NewDecoder(req.Body).Decode(&arithmeticRequest)
	if err != nil {
		return nil, err
	}
	return arithmeticRequest,nil
}

//Response则将对象序列化为JSON格式
func EncodeArithmeticResponse(ctx context.Context, w http.ResponseWriter, response interface{})error{
	//此处如果不设置Content-Type           则value默认为text/plain; charset=utf-8
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

//kit/transport/http包中的NewServer函数实现了http.handler接口，并且装饰了endpoint
//gin.WrapH将http.Handler再装饰成为gin.HandlerFunc中间件
func MakeHttpHandler(srv service.ArithmeticService)http.Handler{
	r := gin.Default()
	arithmeticHandler := httptransport.NewServer(
		endpoint.MakeArithmeticEndpoint(srv),
		DecodeArithmeticRequest,
		EncodeArithmeticResponse,
	)
	r.POST("/add",gin.WrapH(arithmeticHandler))
	return r
}
