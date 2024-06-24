package config

import (
	"fmt"

	goora "github.com/sijms/go-ora/v2"
)

const (
	Mysqldb = "root:123456@tcp(127.0.0.1:3306)/ranking?charset=utf8"
)

var Dsn string

func init() {
	Dsn = goora.BuildUrl("10.251.16.185", 1521, "histdb", "SUTIE", "5Jxz6T^6$", nil)
	fmt.Println(Dsn)
}
