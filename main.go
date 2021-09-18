package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/ilhamtubagus/newsTags/app"
	"github.com/ilhamtubagus/newsTags/infrastructure/cache"
	"github.com/ilhamtubagus/newsTags/infrastructure/persistence"
	"github.com/ilhamtubagus/newsTags/interface/handler"
	"github.com/ilhamtubagus/newsTags/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var dbClient *gorm.DB
var rdbClient *redis.Client

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	// Postgres initialization
	dbClient = utils.GetDatabaseClient()
	// Redis initialization
	rdbClient = utils.GetRedisClient()

}
func main() {
	// Instantiate postgres repositories
	databaseServices := persistence.NewDatabaseRepositories(dbClient)
	err := databaseServices.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}
	// Instantiate redis cacher
	redis := cache.NewRedisCacher(rdbClient)
	// Instantiate apps
	tagApp := app.TagAppImpl{TagRepo: databaseServices.TagRepository}
	topicApp := app.TopicAppImpl{TopicRepo: databaseServices.TopicRepository}
	newsApp := app.NewsAppImpl{NewsRepo: databaseServices.NewsRepository, TagRepo: databaseServices.TagRepository, TopicRepo: databaseServices.TopicRepository}
	// Glue apps with handlers
	tagHandler := handler.NewTagHandler(&tagApp, redis)
	topicHandler := handler.NewTopicHandler(&topicApp, redis)
	newsHandler := handler.NewNewsHandler(&newsApp, redis)

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
