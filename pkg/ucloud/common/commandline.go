/*=============================================================================
#     FileName: commandline.go
#         Desc: parse commandline options
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2016-01-15 20:17:50
#      History:
=============================================================================*/
package ufcommon

import (
	"errors"
	"flag"
	"fmt"
	"reflect"
	"time"
)

var (
	errAddOption = errors.New("Invalid type of option")
)

func AddOption(name string, value interface{}, usage string) (err error) {
	switch v := value.(type) {
	case bool:
		flag.Bool(name, v, usage)
	case string:
		flag.String(name, v, usage)
	case int:
		flag.Int(name, v, usage)
	case int64:
		flag.Int64(name, v, usage)
	case uint:
		flag.Uint(name, v, usage)
	case uint64:
		flag.Uint64(name, v, usage)
	case time.Duration:
		flag.Duration(name, v, usage)
	case float64:
		flag.Float64(name, v, usage)
	default:
		err = errAddOption
	}
	return
}

func getOption(name string) (fg *flag.Flag, err error) {
	fg = flag.Lookup(name)
	if fg == nil {
		err = fmt.Errorf("Can't find option[%s]", name)
	}
	return
}

func GetOptionValue(name string) (value interface{}, err error) {
	fg, err := getOption(name)
	if err != nil {
		return
	}
	if v, ok := fg.Value.(flag.Getter); ok {
		value = v.Get()
	}
	return
}

func GetOptionBoolValue(name string) (value bool, err error) {
	fg, err := getOption(name)
	if err != nil {
		return
	}
	if v, ok := fg.Value.(flag.Getter); ok {
		if va, ok := v.Get().(bool); ok {
			value = va
		} else {
			err = fmt.Errorf("Option [%s]'s type is %v, not bool", name, reflect.TypeOf(v.Get()))
		}
	} else {
		err = fmt.Errorf("Option [%s]'s value haven't implemented Getter Interface", name)
	}
	return
}

func GetOptionStringValue(name string) (value string, err error) {
	fg, err := getOption(name)
	if err != nil {
		return
	}
	if v, ok := fg.Value.(flag.Getter); ok {
		if va, ok := v.Get().(string); ok {
			value = va
		} else {
			err = fmt.Errorf("Option [%s]'s type is %v, not string", name, reflect.TypeOf(v.Get()))
		}
	} else {
		err = fmt.Errorf("Option [%s]'s value haven't implement Getter Interface", name)
	}
	return
}

func GetOptionIntValue(name string) (value int, err error) {
	fg, err := getOption(name)
	if err != nil {
		return
	}
	if v, ok := fg.Value.(flag.Getter); ok {
		if va, ok := v.Get().(int); ok {
			value = va
		} else {
			err = fmt.Errorf("Option [%s]'s type is %v, not int", name, reflect.TypeOf(v.Get()))
		}
	} else {
		err = fmt.Errorf("Option [%s]'s value haven't implement Getter Interface", name)
	}
	return
}

func GetOptionInt64Value(name string) (value int64, err error) {
	fg, err := getOption(name)
	if err != nil {
		return
	}
	if v, ok := fg.Value.(flag.Getter); ok {
		if va, ok := v.Get().(int64); ok {
			value = va
		} else {
			err = fmt.Errorf("Option [%s]'s type is %v, not int64", name, reflect.TypeOf(v.Get()))
		}
	} else {
		err = fmt.Errorf("Option [%s]'s value haven't implement Getter Interface", name)
	}
	return
}

func GetOptionUintValue(name string) (value uint, err error) {
	fg, err := getOption(name)
	if err != nil {
		return
	}
	if v, ok := fg.Value.(flag.Getter); ok {
		if va, ok := v.Get().(uint); ok {
			value = va
		} else {
			err = fmt.Errorf("Option [%s]'s type is %v, not uint", name, reflect.TypeOf(v.Get()))
		}
	} else {
		err = fmt.Errorf("Option [%s] 's value haven't implement Getter Interface", name)
	}
	return
}

func GetOptionUint64Value(name string) (value uint64, err error) {
	fg, err := getOption(name)
	if err != nil {
		return
	}
	if v, ok := fg.Value.(flag.Getter); ok {
		if va, ok := v.Get().(uint64); ok {
			value = va
		} else {
			err = fmt.Errorf("Option [%s]'s type is %v, not uint64", name, reflect.TypeOf(v.Get()))
		}
	} else {
		err = fmt.Errorf("Option [%s]'s value haven't implement Getter Interface", name)
	}
	return
}

func GetOptionFloat64Value(name string) (value float64, err error) {
	fg, err := getOption(name)
	if err != nil {
		return
	}
	if v, ok := fg.Value.(flag.Getter); ok {
		if va, ok := v.Get().(float64); ok {
			value = va
		} else {
			err = fmt.Errorf("Option [%s]'s type is %v, not float64", name, reflect.TypeOf(v.Get()))
		}
	} else {
		err = fmt.Errorf("Option [%s]'s value haven't implement Getter Interface", name)
	}
	return
}

func Usage(err error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	flag.Usage()
}

func ProcessOptions() {
	if !flag.Parsed() {
		flag.Parse()
	}
}

func DumpOptions() {
	if !flag.Parsed() {
		flag.Parse()
	}
	visitor := func(f *flag.Flag) {
		fmt.Println("option =", f.Name, " value =", f.Value)
	}
	flag.VisitAll(visitor)
}
