/*=============================================================================
#     FileName: goid.go
#         Desc: get goroutine id
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2016-04-10 20:17:28
#      History:
=============================================================================*/
package ufcommon

import (
	"bytes"
	"runtime"
	"strconv"
)

func GetGID() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	gid, _ = strconv.ParseUint(string(b), 10, 64)
	return
}
