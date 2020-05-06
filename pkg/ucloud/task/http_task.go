/*=============================================================================
#     FileName: http_task.go
#         Desc: http task
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2016-02-21 20:21:47
#      History:
=============================================================================*/
package uftask

import (
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
	"uframework/common"
)

var (
	httpTaskPoolMu sync.Mutex
	httpTaskPool   = make(map[int32]*HTTPTask)
)

type HTTPTask struct {
	Id         int32
	Gid        uint64
	Pattern    string
	Handler    http.Handler
	FuncName   string
	State      taskState
	timeOut    time.Duration
	isFinished chan bool
}

func NewHTTPTask(pattern string) (task *HTTPTask, err error) {
	taskHandle, err := GetHTTPTaskHandle(pattern)
	if err != nil {
		return
	}
	task = &HTTPTask{
		Id:         atomic.AddInt32(&globalTaskId, 1),
		Pattern:    pattern,
		Handler:    taskHandle.handler,
		State:      stateNew,
		timeOut:    taskHandle.timeOut,
		isFinished: make(chan bool),
	}
	httpTaskPoolMu.Lock()
	httpTaskPool[task.Id] = task
	httpTaskPoolMu.Unlock()
	return
}

func (t *HTTPTask) Run(rw http.ResponseWriter, r *http.Request) (res []byte, err error) {
	t.setState(stateRun)

	funcName := GetTaskFuncName(t.Handler)
	go func() {
		t.Gid = ufcommon.GetGID()
		t.FuncName = funcName
		if ufcommon.CheckWrapPanic() {
			ufcommon.GoSafeHTTP(rw, r, t.Handler.ServeHTTP)
		} else {
			t.Handler.ServeHTTP(rw, r)
		}
		t.isFinished <- true
	}()

	if t.timeOut > 0 {
		select {
		case <-t.isFinished:
			t.setState(stateFinished)
		case <-time.After(t.timeOut):
			t.setState(stateFinished)
			err = errors.New("task timet out")

		}
	} else {
		select {
		case <-t.isFinished:
			t.setState(stateFinished)
		}
	}

	httpTaskPoolMu.Lock()
	delete(httpTaskPool, t.Id)
	httpTaskPoolMu.Unlock()
	return
}

func (t *HTTPTask) setState(state taskState) {
	t.State = state
}

func LenHTTPTasks() int {
	return len(httpTaskPool)
}

func GetHTTPTaskByGid(gid uint64) (task interface{}) {
	httpTaskPoolMu.Lock()
	defer httpTaskPoolMu.Unlock()
	for _, t := range httpTaskPool {
		if t.Gid == gid {
			return t
		}
	}
	return nil
}

func DumpHTTPTasks() (tasks map[int32]*HTTPTask) {
	tasks = make(map[int32]*HTTPTask)
	httpTaskPoolMu.Lock()
	defer httpTaskPoolMu.Unlock()
	for k, v := range httpTaskPool {
		tasks[k] = v
	}
	return
}
