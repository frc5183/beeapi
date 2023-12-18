package database

import (
	"beeapi/config"
	"beeapi/models"
	mysql2 "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var db *gorm.DB

var mdls = []interface{}{
	models.Checkout{},
	models.Item{},
	models.User{},
}

func Init() {
	cfg := mysql2.NewConfig()
	cfg.User = config.GetConfig().Database.User
	cfg.Passwd = config.GetConfig().Database.Password
	cfg.Net = "tcp"
	cfg.Addr = config.GetConfig().Database.Host + ":" + strconv.Itoa(config.GetConfig().Database.Port)
	cfg.DBName = config.GetConfig().Database.Name
	cfg.ParseTime = true
	cfg.MultiStatements = true

	var err error

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSNConfig: cfg,
	}))

	if err != nil {
		panic(err)
	}

	// TODO: will crash if database is not created yet

	migrate()
}

// TODO: Add custom migrations system
func migrate() {
	for _, model := range mdls {
		err := db.AutoMigrate(model)
		if err != nil {
			panic(err)
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
