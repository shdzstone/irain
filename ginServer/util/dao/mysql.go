package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"

	"irain/ginServer/config"
	"log"
)

func DB() *gorm.DB {

	cfg := config.MysqlConf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	log.Println(dsn)
	
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "bl_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,  // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	log.Println(dsn)

	if err != nil {
		log.Println("database open err:", err)
	}
	//db默认设置
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(500)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
