/*=============================================================================
#     FileName: http_server.go
#         Desc: http server
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2016-01-20 20:24:53
#      History:
=============================================================================*/
package ufnet

import (
	"net/http"

	"apiServer/ucloud/net/websocket"
)

func listenAndServeHTTP(addr string) error {
	http.HandleFunc("/", RouteHTTP)
	http.Handle("/ws", websocket.Handler(RouteWs))
	return http.ListenAndServe(addr, nil)
}

/*
	自定义multiplexer，使用方式
	mux中可包括:
	mux.HandleFunc("/", RouteHTTP)
	mux.Handle("/ws", websocket.Handler(RouteWs))
*/
func listenAndServeHTTPMux(addr string, mux http.Handler) error {
	return http.ListenAndServe(addr, mux)
}
