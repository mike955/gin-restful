/*=============================================================================
#     FileName: http_common.go
#         Desc: http common
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2016-01-20 20:24:40
#      History:
=============================================================================*/
package ufnet

import (
	"net/http"

	"apiServer/ucloud/net/websocket"
)

var (
	RouteHTTP = func(w http.ResponseWriter, r *http.Request) {}
	RouteWs   = func(ws *websocket.Conn) {}
)
