package http

import (
	"LTest/src/routes"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mgutz/logxi/v1"
	"github.com/valyala/fasthttp"
)

var (
	listener net.Listener
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	path := strings.Split(string(ctx.Path()), "/")

	switch path[1] {
	case "users":
		routes.HandleUsers(ctx)
	default:
		ctx.Error("Handler not found", fasthttp.StatusNotFound)
	}
}

func StartHttp(port int) {
	go func() {
		var err error
		listener, err = net.Listen("tcp4", "0.0.0.0:"+strconv.Itoa(port))
		if err != nil {
			log.Info("HTTP: cant create listener, wait and try")
			ticker := time.NewTicker(time.Millisecond * 100)
			for range ticker.C {
				listener, err = net.Listen("tcp4", "0.0.0.0:"+strconv.Itoa(port))
				if err == nil {
					break
				}
			}
		}

		log.Info("HTTP: wait connections on 0.0.0.0")
		s := &fasthttp.Server{
			Name:               "LTest",
			Handler:            requestHandler,
			ReadBufferSize:     1024,
			IdleTimeout:        30 * time.Second,
			TCPKeepalivePeriod: time.Minute,
			DisableKeepalive:   false,
			TCPKeepalive:       true,
		}

		err = s.Serve(listener)
		if err != nil {
			log.Fatal("HTTP: can't run http server", err)
			os.Exit(1)
		}
	}()
}
