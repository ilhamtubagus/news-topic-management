package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ilhamtubagus/newsTags/app"
	"github.com/ilhamtubagus/newsTags/infrastructure/persistence"
	"github.com/ilhamtubagus/newsTags/interface/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	DSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	postgresClient, err := gorm.Open(postgres.New(postgres.Config{DSN: DSN}))
	if err != nil {
		log.Fatal(err)
	}
	dbClient = postgresClient

}
func main() {
	// Instantiate database repositories
	databaseServices := persistence.NewDatabaseRepositories(dbClient)
	err := databaseServices.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}
	// Instantiate Apps
	tagApp := app.TagAppImpl{TagRepo: databaseServices.TagRepository}
	topicApp := app.TopicAppImpl{TopicRepo: databaseServices.TopicRepository}
	newsApp := app.NewsAppImpl{NewsRepo: databaseServices.NewsRepository, TagRepo: databaseServices.TagRepository, TopicRepo: databaseServices.TopicRepository}
	// Glue apps with handlers
	tagHandler := handler.NewTagHandler(&tagApp)
	topicHandler := handler.NewTopicHandler(&topicApp)
	newsHandler := handler.NewNewsandler(&newsApp)

	// Define routes
	e := echo.New()
	// Tag routes
	e.POST("/tag", tagHandler.SaveTag)
	e.PUT("/tag/:id", tagHandler.UpdateTag)
	e.GET("/tag/:id", tagHandler.GetTagById)
	e.DELETE("/tag/:id", tagHandler.DeleteTag)
	e.GET("/tag", tagHandler.GetAllTag)
	// Topic routes
	e.POST("/topic", topicHandler.SaveTopic)
	e.PUT("/topic/:id", topicHandler.UpdateTopic)
	e.GET("/topic/:id", topicHandler.GetTopicById)
	e.DELETE("/topic/:id", topicHandler.DeleteTopic)
	e.GET("/topic", topicHandler.GetAllTopic)
	// News Routes
	e.POST("/news", newsHandler.SaveNews)
	e.GET("/news/:id", newsHandler.GetNewsById)
	e.GET("/news", newsHandler.GetAllNews)
	e.PUT("/news/:id", newsHandler.UpdateNews)
	e.DELETE("/news/:id", newsHandler.DeleteNews)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
