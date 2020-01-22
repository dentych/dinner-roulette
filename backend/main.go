package main

import (
	"fmt"
	"github.com/dentych/dinner-dash/config"
	"github.com/dentych/dinner-dash/handlers"
	"github.com/dentych/dinner-dash/database"
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/middleware"
	"github.com/dentych/dinner-dash/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	configuration := config.FromEnv()

	logging.Init()
	database.Init(configuration.DbConfig)
	database.RunMigrations(configuration.DbConfig)

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{"http://localhost:8080", "http://dinner-dash.tychsen.me", "https://dinner-dash.tychsen.me"}
	//corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"authorization", "content-type", "charset"}
	corsConfig.AllowWildcard = true
	router.Use(cors.New(corsConfig))

	recipeDao := database.RecipeDao{}
	userDao := database.UserDao{}

	authController := handlers.NewAuthController(userDao, configuration.CookieHost)
	recipeHandler := handlers.NewRecipeHandler(recipeDao)

	// Provide frontend files
	router.StaticFile("/", "dist/index.html")
	router.StaticFile("/favicon.ico", "dist/favicon.ico")
	router.Static("/assets", "dist/assets")
	router.NoRoute(func(c *gin.Context) {
		if strings.Index(c.Request.RequestURI, "api") > -1 {
			c.JSON(400, "Bad requesteroo")
			return
		}
		c.File("../frontend/dist/index.html")
	})

	unprotectedApiRouter := router.Group("/api")
	unprotectedApiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})

	unprotectedApiRouter.POST("/login", authController.Login)
	unprotectedApiRouter.POST("/register", authController.Register)
	unprotectedApiRouter.POST("/token", authController.Token)
	unprotectedApiRouter.POST("/logout", authController.Logout)

	protectedApiRouter := router.Group("/api", middleware.AuthRequired)
	protectedApiRouter.GET("/test", func(c *gin.Context) {
		userid := c.GetInt("userid")
		c.JSON(200, fmt.Sprintf("Authenticated as: %v", userid))
	})

	protectedApiRouter.POST("/recipes", recipeHandler.CreateRecipe)

	protectedApiRouter.GET("/recipes", recipeHandler.GetRecipes)

	protectedApiRouter.GET("/recipes/:id", recipeHandler.GetRecipeById)

	protectedApiRouter.PUT("/recipes/:id", recipeHandler.UpdateRecipe)

	protectedApiRouter.DELETE("/recipes/:id", recipeHandler.DeleteRecipe)

	protectedApiRouter.POST("/mealplan", func(c *gin.Context) {
		recipes, err := recipeDao.GetAll(c.GetInt("uid"))
		if err != nil {
			c.JSON(500, "Error getting recipes while generating meal plan")
			return
		}

		if len(recipes) < 1 {
			c.JSON(400, "Can't generate meal plan - there are no recipes for the current user")
			return
		}

		var body gin.H
		err = c.BindJSON(&body)
		if err != nil {
			c.JSON(500, "error")
			logging.Error.Println("Error while reading body of /mealplan", err)
			return
		}

		mealplan := make([]models.Recipe, 0, 7)

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 7; i++ {
			index := r.Intn(len(recipes))
			mealplan = append(mealplan, recipes[index])
		}

		c.JSON(200, mealplan)
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Such error\n")
	}
}
