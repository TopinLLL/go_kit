package main

import (
	"go_kit/service"
	"go_kit/transport"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

func main()  {
	var logger  = log.NewLogfmtLogger(os.Stderr)
	var srv  =service.ArithmeticService{}
	var h = transport.MakeHttpHandler(srv)

	logger.Log("msg", "HTTP", "addr", ":9999")
	logger.Log("err", http.ListenAndServe(":9999", h))
}
