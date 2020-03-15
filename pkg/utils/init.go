package utils

import "gin-restful/pkg/utils/library"

func Setup()  {
	library.MysqlSetup()
	library.RedisSetup()
}
