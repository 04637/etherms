package internal

import (
	"aid.dev/etherms/config"
	"aid.dev/etherms/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

// InitDB 根据配置文件初始化数据库
func InitDB(conf config.Conf) {
	dsn := conf.DB.Username + ":" + conf.DB.Password + "@tcp(" + conf.DB.Host + ")/etherpad?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	HandleErr(err)
	sqlDB, err := DB.DB()
	HandleErr(err)
	sqlDB.SetMaxIdleConns(conf.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.DB.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(conf.DB.ConnMaxLifetime))
	migrateSchema()
}

func migrateSchema() {
	err := DB.AutoMigrate(&model.User{})
	HandleErr(err)
}
