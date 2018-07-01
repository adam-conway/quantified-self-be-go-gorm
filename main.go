package main
import (
 "fmt"
 "github.com/gin-gonic/gin"
 "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/postgres"
)
var db *gorm.DB
var err error

type Food struct {
 ID uint `json:"id"`
 Name string `json:"name"`
 Calories int `json:"calories"`
}


func main() {
 db, err = gorm.Open("postgres", "host=localhost port=5432 user=adamconway dbname=adamconway sslmode=disable")
 if err != nil {
 fmt.Println(err)
 }
 defer db.Close()
db.AutoMigrate(&Food{})
 r := gin.Default()
 r.GET("/api/v1/foods", GetFoods)
 r.GET("/api/v1/foods/:id", GetFood)
 r.POST("/api/v1/foods", CreateFood)
r.Run(":8080")
}

func CreateFood(c *gin.Context) {
 var food Food
 c.BindJSON(&food)
 db.Create(&food)
 c.JSON(200, food)
}

func GetFood(c *gin.Context) {
 id := c.Params.ByName("id")
 var food Food
 if err := db.Where("id = ?", id).First(&food).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, food)
 }
}

func GetFoods(c *gin.Context) {
 var foods []Food
 if err := db.Find(&foods).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, foods)
 }
}
