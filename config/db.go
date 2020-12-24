package config

import(
	"../models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)



var DB *gorm.DB

func InitDB() {
	var err error

	dsn := "fajar:fajar@tcp(localhost:3306)/learn_gin?charset=utf8mb4&parseTime=True&loc=Local"
  	DB,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}
	
	// defer db.Close()

	//Migrate Schema
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Article{})

	// DB.AutoMigrate(&models.Article{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	// DB.Model(&models.User{}).Related(&models.Article{})

	
}