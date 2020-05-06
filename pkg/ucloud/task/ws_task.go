/*=============================================================================
#     FileName: ws_task.go
#         Desc: web socket task
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2016-04-10 20:19:38
#      History:
=============================================================================*/
package uftask

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
	"uframework/common"
)

var (
	wsTaskPoolMu sync.Mutex
	wsTaskPool   = make(map[int32]*WsTask)
)

type WsTask struct {
	Id       int32
	Gid      uint64
	Pattern  string
	Handler  WsTaskHandler
	FuncName string
	State    taskState
	timeOut  time.Duration
	msgChan  chan []byte
}

func NewWsTask(pattern string) (task *WsTask, err error) {
	taskHandle, err := GetWsTaskHandle(pattern)
	if err != nil {
		return
	}
	task = &WsTask{
		Id:      atomic.AddInt32(&globalTaskId, 1),
		Pattern: pattern,
		Handler: taskHandle.handler,
		State:   stateNew,
		timeOut: taskHandle.timeOut,
		msgChan: make(chan []byte),
	}
	wsTaskPoolMu.Lock()
	wsTaskPool[task.Id] = task
	wsTaskPoolMu.Unlock()
	return
}

func (t *WsTask) Run(req interface{}, conn interface{}) (res []byte, err error) {
	t.setState(stateRun)
	funcName := GetTaskFuncName(t.Handler)
	var ok bool
	go func() {
		t.Gid = ufcommon.GetGID()
		t.FuncName = funcName
		t.Handler.ServeWs(t.msgChan, req, conn)
	}()

	if t.timeOut > 0 {
		select {
		case res, ok = <-t.msgChan:
			if !ok {
				err = errors.New("task fail ,close")
			}
			t.setState(stateFinished)
		case <-time.After(t.timeOut):
			t.setState(stateFinished)
			err = errors.New("task timet out")

		}
	} else {
		select {
		case res, ok = <-t.msgChan:
			if !ok {
				err = errors.New("task fail ,close")
			}
			t.setState(stateFinished)
		}
	}
	wsTaskPoolMu.Lock()
	delete(wsTaskPool, t.Id)
	wsTaskPoolMu.Unlock()
	return
}

func (t *WsTask) setState(state taskState) {
	t.State = state
}

func LenWsTasks() int {
	return len(wsTaskPool)
}

func GetWsTaskByGid(gid uint64) (task interface{}) {
	wsTaskPoolMu.Lock()
	defer wsTaskPoolMu.Unlock()
	for _, t := range wsTaskPool {
		if t.Gid == gid {
			return t
		}
	}
	return nil
}

func DumpWsTasks() (tasks map[int32]*WsTask) {
	tasks = make(map[int32]*WsTask)
	wsTaskPoolMu.Lock()
	defer wsTaskPoolMu.Unlock()
	for k, v := range wsTaskPool {
		tasks[k] = v
	}
	return
}
