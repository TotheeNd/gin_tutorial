package dao

import (
	"gin-ranking/config"
	"gin-ranking/pkg/logger"

	// "time"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"database/sql"

	"fmt"
	"log"

	goora "github.com/sijms/go-ora/v2"
)

var (
	DB  *sql.DB
	err error
)

func init() {

	fmt.Println("程序执行进入DAO包中!")

	// Register the go-ora driver
	sql.Register("goora", goora.NewDriver())

	// 创建数据库连接
	DB, err = sql.Open("goora", config.Dsn)

	// Db, err = gorm.Open("mysql", config.Mysqldb)
	if err != nil {
		fmt.Println("数据库连接失败")
		logger.Error(map[string]interface{}{"oracle connect error": err.Error()})
	}
	fmt.Println("数据库连接成功!")
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// if Db.Error != nil {
	// 	logger.Error(map[string]interface{}{"database error": Db.Error})
	// }

	// // ----------------------- 连接池设置 -----------------------
	// // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	// Db.DB().SetMaxIdleConns(10)

	// // SetMaxOpenConns 设置打开数据库连接的最大数量。
	// Db.DB().SetMaxOpenConns(100)

	// // SetConnMaxLifetime 设置了连接可复用的最大时间。
	// Db.DB().SetConnMaxLifetime(time.Hour)
}
