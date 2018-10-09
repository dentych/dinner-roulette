package main

import (
	"fmt"
	"github.com/dentych/dinner-dash/database"
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/middleware"
	"github.com/dentych/dinner-dash/model"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	logging.Init()
	database.Init()

	router := gin.Default()

	recipeDao := database.RecipeDao{}
	ingredientDao := database.IngredientDao{}

	unprotectedApiRouter := router.Group("/api")
	unprotectedApiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})

	protectedApiRouter := router.Group("/api", middleware.AuthRequired())
	protectedApiRouter.GET("/john", func(c *gin.Context) {
		user := c.GetString("user")
		c.JSON(200, fmt.Sprintf("Authenticated as: %s", user))
	})

	// Recipe stuff
	protectedApiRouter.GET("/recipes", func(c *gin.Context) {
		recipes, err := recipeDao.GetAll(c.GetString("user"))
		if err != nil {
			c.JSON(500, "error while getting recipes")
			return
		}

		c.JSON(200, recipes)
	})
	protectedApiRouter.GET("/recipes/:id", func(c *gin.Context) {
		id := c.Param("id")

		recipe := recipeDao.GetById(c.GetString("user"), id, false)

		if recipe != nil {
			c.JSON(200, *recipe)
		} else {
			c.JSON(404, "not found")
		}
	})
	protectedApiRouter.PUT("/recipes", func(c *gin.Context) {
		user := c.GetString("user")
		var recipe model.Recipe
		err := c.BindJSON(&recipe)
		if err != nil {
			logging.Error.Printf("Could not parse recipe: %s", err)
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
	protectedApiRouter.POST("/recipes", func(c *gin.Context) {
		var recipe model.Recipe
		err := c.BindJSON(&recipe)
		if err != nil {
			logging.Error.Printf("Error: %s", err)
			return
		}

		if err = validateRecipe(recipe, false); err != nil {
			logging.Error.Printf("Error validating recipe: %s", err)
			c.JSON(400, err.Error())
			return
		}

		if err = recipeDao.Insert(c.GetString("user"), &recipe); err != nil {
			logging.Error.Printf("Error creating recipe: %s", err)
			c.JSON(500, "Error while creating recipe.")
		}

		c.JSON(201, recipe)
	})
	protectedApiRouter.POST("/recipes/:id/ingredients", func(c *gin.Context) {
		var ingredient model.Ingredient
		if err := c.BindJSON(&ingredient); err != nil {
			logging.Error.Printf("Error: %s", err)
			return
		}

		recipe := recipeDao.GetById(c.GetString("user"), c.Param("id"), false)
		if recipe == nil {
			c.JSON(400, "recipe not found")
			return
		}

		if err := recipeDao.AddIngredient(c.Param("id"), ingredient); err != nil {
			logging.Error.Printf("Error while adding ingredient to recipe: %s", err)
			c.JSON(400, "error while adding ingredient")
			return
		}

		c.Status(201)
	})
	protectedApiRouter.DELETE("/recipes/:id", func(c *gin.Context) {
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

	// Ingredients
	protectedApiRouter.GET("/ingredients", func(c *gin.Context) {
		ingredientDao.GetAll()
	})
	protectedApiRouter.POST("/ingredients", func(c *gin.Context) {
		var ingredient model.Ingredient
		if err := c.BindJSON(&ingredient); err != nil {
			logging.Error.Printf("Error while parsing ingredient: %s", err)
			return
		}

		if err := ingredientDao.Insert(&ingredient); err != nil {
			logging.Error.Printf("Error while inserting ingredient in database: %s", err)
			c.JSON(500, "Error inserting ingredient in database")
			return
		}

		c.Status(201)
	})

	// Mealplans
	protectedApiRouter.POST("/mealplans", func(c *gin.Context) {
		recipes, err := recipeDao.GetAll(c.GetString("user"))
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
			logging.Error.Println("Error while parsing meal plan", err)
			return
		}

		mealPlan := make([]model.Recipe, 0, 7)

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 7; i++ {
			index := r.Intn(len(recipes))
			mealPlan = append(mealPlan, recipes[index])
		}

		c.JSON(201, mealPlan)
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
