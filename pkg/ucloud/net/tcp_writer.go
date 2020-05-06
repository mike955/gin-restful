/*=============================================================================
#     FileName: tcp_writer.go
#         Desc: tcp server conneciton writer
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2016-01-20 20:26:56
#      History:
=============================================================================*/
package ufnet

import (
	"encoding/binary"
	"io"
)

type writer struct {
	w   io.Writer
	buf []byte
}

func newWriter(w io.Writer) *writer {
	return &writer{
		w:   w,
		buf: make([]byte, defaultHeadSize),
	}
}

func (w *writer) writePacket(packet []byte) (n int, err error) {
	// 将数据包大小写入头部
	n, err = w.writeHead(len(packet))
	if err != nil {
		return 0, err
	}
	// 发送数据包
	return w.writeBody(packet)
}

func (w *writer) writeHead(plen int) (n int, err error) {
	return w.writeUint32BE(uint32(plen))
}

func (w *writer) writeUint32BE(v uint32) (n int, err error) {
	binary.BigEndian.PutUint32(w.buf[:defaultHeadSize], v)
	return w.w.Write(w.buf[:defaultHeadSize])
}

func (w *writer) writeBody(body []byte) (n int, err error) {
	return w.w.Write(body)
}
