package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB = mysqlConnection()

func mysqlConnection() *gorm.DB {
	// migrate -database "mysql://root:Arifin123!@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations down
	// migrate -database 'mysql://root:Arifin123!@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local' -path db/migrations down
	dsn := "root:Arifin123!@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	file, err := os.Create("gorm-log.txt")
	if err != nil {
		// Handle error
		log.Println(err.Error())
	}
	newLogger := logger.New(
		log.New(file, "", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,         // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Println(err.Error())
	}
	return db
}
