package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var DB *gorm.DB

func initDatabase() {
	var err error
	dsn := "host=postgres user=postgres password=postgres dbname=todos port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Todo{})
}

func main() {
	initDatabase()

	r := gin.Default()

	// CORS middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:" + os.Getenv("ALLOWED_PORT")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		DB.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.Create(&todo)
		c.JSON(http.StatusOK, todo)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var todo Todo
		if err := DB.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		var updatedTodo Todo
		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo.Completed = updatedTodo.Completed
		err = DB.Save(&todo).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update"})
			return
		}

		c.JSON(http.StatusOK, todo)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		DB.Delete(&Todo{}, id)
		c.Status(http.StatusNoContent)
	})

	r.Run(":8080")
}
