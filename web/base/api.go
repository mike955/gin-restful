package base

import (
	"path"
	"regexp"
	"runtime"
	"strings"
)

type Api struct {
	Run func(args *ApiArgs) interface{}
}

var apis map[string]Api

func fileName() (pkg, file string) {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[1])
	file, _ = f.FileLine(pc[1])
	var dir string
	dir, file = path.Split(file)
	pkg = path.Base(dir)
	return pkg, file
}

func NewApi(run func(args *ApiArgs) interface{}) {
	_, file := fileName()
	reg := regexp.MustCompile("^(.+)\\.go$")
	matched := reg.FindAllStringSubmatch(file, -1)
	if len(matched) == 0 {
		panic("Invalid file name must be api_xxx.go")
	}
	apiName := strings.Replace(matched[0][1], "_", ".", -1)
	if apis == nil {
		apis = make(map[string]Api)
	}
	apis[apiName] = Api{
		Run: run,
	}
}

func ExistsAPI(action string) bool {
	_, ok := apis[action]
	if ok {
		return true
	}
	return false
}

func Call(args *ApiArgs) interface{}{
	return apis[args.Name].Run(args)
}