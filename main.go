package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gosimple/slug"
	//   "fmt"

)

type Article struct {
	gorm.Model
	Title string
	Slug string `gorm:"unique_index"`
	Desc string `sql:"type:text;"`

}

var DB *gorm.DB

func main() {
	var err error

	dsn := "fajar:fajar@tcp(localhost:3306)/learn_gin?charset=utf8mb4&parseTime=True&loc=Local"
  	DB,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}
	
	// defer db.Close()

	//Migrate Schema
	DB.AutoMigrate(&Article{})

	router := gin.Default()

	v1 := router.Group("api/")
	{

		article := v1.Group("/article")
		{
			
			article.GET("/", getHome)
			article.GET("/:slug", getArticle)
			article.POST("/", postArticle)
		}
	}

	router.Run()

	
}

func getHome(c *gin.Context){

	items := []Article{}
	DB.Find(&items)


	c.JSON(200, gin.H{
		"status":"berhasil",
		"data": items,
})
}

func getArticle(c *gin.Context){
	slug := c.Param("slug")
	
	var item Article

	// v_first := DB.First(&item, "slug = ?", slug).Error
	// errors.Is(v_first, ErrRecordNotFound)

	if result := DB.First(&item, "slug = ?", slug); result.Error != nil{
		c.JSON(404, gin.H{"status":"error", "message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data": item,
	})
}

func postArticle(c *gin.Context){

	item := Article{
	Title : c.PostForm("title"),
	Slug  : slug.Make(c.PostForm("title")),
	Desc : c.PostForm("desc"),
	}

	DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data": item,
	})
}

