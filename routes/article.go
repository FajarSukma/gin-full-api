package routes

import(
	"fmt"
	"time"
	"strconv"
	"../config"
	"../models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"reflect"
)


 

func GetHome(c *gin.Context){

	items := []models.Article{}
	config.DB.Find(&items)


	c.JSON(200, gin.H{
		"status":"berhasil",
		"data": items,
})
}

func GetArticle(c *gin.Context){
	slug := c.Param("slug")
	
	var item models.Article

	// v_first := DB.First(&item, "slug = ?", slug).Error
	// errors.Is(v_first, ErrRecordNotFound)

	if result := config.DB.First(&item, "slug = ?", slug); result.Error != nil{
		c.JSON(404, gin.H{"status":"error", "message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data": item,
	})
}

func PostArticle(c *gin.Context){

	var oldItem models.Article
	slug := slug.Make(c.PostForm("title"))
	
	if result := config.DB.First(&oldItem, "slug = ?", slug); result.Error == nil{
		slug = slug + strconv.FormatInt(time.Now().Unix(), 10)
	}

	item := models.Article{
		Title : c.PostForm("title"),
		Slug  : slug,
		Desc : c.PostForm("desc"),
		Tag : c.PostForm("tag"),
		UserID: uint(c.MustGet("jwt_user_id").(float64)),
	}

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data": item,
	})
}


func GetArticleByTag(c *gin.Context){
	 tag := c.Param("tag")
	 items :=[]models.Article{}

	//  x := config.DB.Where("tag LIKE ? ", "%" + tag + "%").Find(&items)


	 if result := config.DB.Where("tag LIKE ? ", "%" + tag + "%").Find(&items); result == nil {
		fmt.Println("Get Article if")
		 c.JSON(404, gin.H{"status":"error", "message": "record not found"})
		 c.Abort()
		 return
	}else{

	 c.JSON(200, gin.H{"data": items})
	}
}

func UpdateArticle(c *gin.Context){
	id := c.Param("id")
	
	var item models.Article

	// v_first := DB.First(&item, "slug = ?", slug).Error
	// errors.Is(v_first, ErrRecordNotFound)

	if result := config.DB.First(&item, "id = ?", id); result.Error != nil{
		c.JSON(404, gin.H{"status":"error", "message": "record not found"})
		c.Abort()
		return
	}

	if uint(c.MustGet("jwt_user_id").(float64)) != item.UserID{
		c.JSON(404, gin.H{"status":"error", "message": "invalid user for update data"})
		c.Abort()
		return
	}

	config.DB.Model(&item).Where("id = ?", id).Updates(models.Article{
		Title : c.PostForm("title"),
		Desc : c.PostForm("desc"),
		Tag : c.PostForm("tag"),
	})

	c.JSON(200, gin.H{
		"status": "berhasil update",
		"data": item,
	})
}

func GetProfile(c *gin.Context) {
	fmt.Println("Tracer 1")
	var user models.User
	fmt.Println("Tracer 2")
	user_id := int(c.MustGet("jwt_user_id").(float64))
	fmt.Println("Tracer 3", user_id)
	// item := config.DB.Where("id = ?", user_id).Preload("Articles", "user_id = ?", user_id).Find(&user)
	item := config.DB.Preload("Articles", "user_id = ?", user_id).Find(&user)
	fmt.Println("Tracer 4", item, reflect.TypeOf(item))
	
	
	// _, err := json.Marshal(item)
	// fmt.Println(err)
	
	c.JSON(200, gin.H{
		"data": item,
		"status": "berhasil update",
		
	})
	fmt.Println("Tracer 6")
}


func DeleteArticle (c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	config.DB.Where("id = ?", id).Delete(&article)

	c.JSON(200, gin.H{
		"data": article,
		"status": "berhasil deletee",
		
	})
}