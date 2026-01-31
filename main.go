package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	user_handler "go.learning.com/go2025/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//connect database
	dsn := "host=localhost user=root password=123456 dbname=daithuvien port=5450 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Connectdb successful", db)

	//create gin http server
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	v1Route := r.Group("/api/v1")
	{
		// Define a simple GET endpoint
		v1Route.GET("/healthcheck", func(c *gin.Context) {
			// Return JSON response
			c.JSON(http.StatusOK, gin.H{
				"message": "Server is running now",
			})
		})

		//User
		usersRoute := v1Route.Group("/users")
		{
			usersRoute.GET("/", user_handler.GetAllUsersHandler(db))      //Get All users
			usersRoute.GET("/:id", user_handler.GetUserByIdHandler(db))   //Get User by id
			usersRoute.PATCH("/:id", user_handler.UpdateUserHandler(db))  //Update user by id
			usersRoute.DELETE("/:id", user_handler.DeleteUserHandler(db)) //Delete user by id
			usersRoute.POST("", user_handler.CreatedUserHandler(db))      //Create new user
		}

		storiesRoute := v1Route.Group("/stories")
		{
			storiesRoute.GET("/")       //Get All stories
			storiesRoute.GET("/:id")    //Get stories by id
			storiesRoute.PATCH("/:id")  //Update stories by id
			storiesRoute.DELETE("/:id") //Delete stories by id
			storiesRoute.POST("")       //Create new stories
		}

	}

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
