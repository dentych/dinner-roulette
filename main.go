package main

import (
	"fmt"
	"github.com/dentych/dinner-dash/database"
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/middleware"
	"github.com/dentych/dinner-dash/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

var (
	InfoLog *log.Logger
	ErrorLog *log.Logger
)

func main() {
	logging.Init()
	database.Init()

	router := gin.Default()

	recipeDao := database.RecipeDao{}

	unprotectedApiRouter := router.Group("/api")
	unprotectedApiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})

	protectedApiRouter := router.Group("/api", middleware.AuthRequired())
	protectedApiRouter.GET("/john", func(c *gin.Context) {
		user := c.GetString("user")
		c.JSON(200, fmt.Sprintf("Authenticated as: %s", user))
	})
	protectedApiRouter.PUT("/recipe", func(c *gin.Context) {
		user := c.GetString("user")
		var recipe model.Recipe
		err := c.MustBindWith(&recipe, binding.JSON)
		if err != nil {
			logging.Error.Printf("Could not parse recipe: %s", err)
			c.JSON(400, err.Error())
			return
		}

		err = validateRecipe(recipe, true)
		if err != nil {
			logging.Info.Printf("Could not validate recipe object: %s", err)
			c.JSON(400, "invalid recipe object")
			return
		}

		err = recipeDao.Update(user, recipe)
		if err != nil {
			logging.Error.Printf("Error when updating recipe: %s", err)
			c.JSON(500, "error when updating recipe: "+err.Error())
			return
		}

		c.JSON(200, "updated")
	})
	protectedApiRouter.POST("/recipe", func(c *gin.Context) {
		var recipe model.Recipe
		err := c.MustBindWith(&recipe, binding.JSON)
		if err != nil {
			logging.Error.Printf("Error: %s", err)
			return
		}

		err = validateRecipe(recipe, false)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		recipeDao.Insert(c.GetString("user"), &recipe)

		c.JSON(201, "created")
	})
	protectedApiRouter.GET("/recipe/:id", func(c *gin.Context) {
		idAsString := c.Param("id")
		id, err := strconv.ParseInt(idAsString, 10, 64)
		if err != nil {
			logging.Error.Printf("Error: %s", err)
			c.JSON(400, "ID can't be parsed as int")
			return
		}

		recipe := recipeDao.GetById(c.GetString("user"), id)

		if recipe != nil {
			c.JSON(200, *recipe)
		} else {
			c.JSON(404, "not found")
		}
	})
	protectedApiRouter.DELETE("/recipe/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			logging.Error.Printf("Error: %s", err)
			c.JSON(400, "invalid or missing id missing from URL path (/api/recipe/{id})")
			return
		}

		recipeDeleted, err := recipeDao.Delete(c.GetString("user"), id)
		if err != nil {
			logging.Error.Printf("Error deleting recipe: %s", err)
			c.JSON(500, "error while deleting recipe")
			return
		}
		if recipeDeleted {
			c.JSON(200, "recipe deleted")
		} else {
			c.JSON(404, "recipe not found")
		}
	})

	router.Run(":8080")
}
func validateRecipe(recipe model.Recipe, checkId bool) error {
	if checkId {
		if recipe.ID < 1 {
			return fmt.Errorf("missing ID")
		}
	}
	if len(recipe.Name) < 1 {
		return fmt.Errorf("missing recipe name")
	}

	return nil
}
