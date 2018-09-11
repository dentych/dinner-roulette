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

	mealDao := database.MealDao{}

	unprotectedApiRouter := router.Group("/api")
	unprotectedApiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})

	protectedApiRouter := router.Group("/api", middleware.AuthRequired())
	protectedApiRouter.GET("/john", func(c *gin.Context) {
		user := c.GetString("user")
		c.JSON(200, fmt.Sprintf("Authenticated as: %s", user))
	})
	protectedApiRouter.PUT("/meal", func(c *gin.Context) {
		user := c.GetString("user")
		var meal model.Meal
		err := c.MustBindWith(&meal, binding.JSON)
		if err != nil {
			logging.Error.Printf("Could not parse meal: %s", err)
			c.JSON(400, err.Error())
			return
		}

		err = validateMeal(meal, true)
		if err != nil {
			logging.Info.Printf("Could not validate meal object: %s", err)
			c.JSON(400, "invalid meal object")
			return
		}

		err = mealDao.Update(user, meal)
		if err != nil {
			logging.Error.Printf("Error when updating meal: %s", err)
			c.JSON(500, "error when updating meal: "+err.Error())
			return
		}

		c.JSON(200, "updated")
	})
	protectedApiRouter.POST("/meal", func(c *gin.Context) {
		var meal model.Meal
		err := c.MustBindWith(&meal, binding.JSON)
		if err != nil {
			logging.Error.Printf("Error: %s", err)
			return
		}

		err = validateMeal(meal, false)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		mealDao.Insert(c.GetString("user"), &meal)

		c.JSON(201, "created")
	})
	protectedApiRouter.GET("/meal/:id", func(c *gin.Context) {
		idAsString := c.Param("id")
		id, err := strconv.ParseInt(idAsString, 10, 64)
		if err != nil {
			logging.Error.Printf("Error: %s", err)
			c.JSON(400, "ID can't be parsed as int")
			return
		}

		meal := mealDao.GetById(c.GetString("user"), id)

		if meal != nil {
			c.JSON(200, *meal)
		} else {
			c.JSON(404, "not found")
		}
	})
	protectedApiRouter.DELETE("/meal/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			logging.Error.Printf("Error: %s", err)
			c.JSON(400, "invalid or missing id missing from URL path (/api/meal/{id})")
			return
		}

		mealDeleted, err := mealDao.Delete(c.GetString("user"), id)
		if err != nil {
			logging.Error.Printf("Error deleting meal: %s", err)
			c.JSON(500, "error while deleting meal")
			return
		}
		if mealDeleted {
			c.JSON(200, "meal deleted")
		} else {
			c.JSON(404, "meal not found")
		}
	})

	router.Run(":8080")
}
func validateMeal(meal model.Meal, checkId bool) error {
	if checkId {
		if meal.ID < 1 {
			return fmt.Errorf("missing ID")
		}
	}
	if len(meal.Name) < 1 {
		return fmt.Errorf("missing meal name")
	}

	return nil
}
