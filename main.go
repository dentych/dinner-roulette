package main

import (
	"fmt"
	"github.com/dentych/dinner-dash/database"
	"github.com/dentych/dinner-dash/middleware"
	"github.com/dentych/dinner-dash/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()

	mealDao := database.MealDao{}

	unprotectedApiRouter := router.Group("/api")
	unprotectedApiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})

	protectedApiRouter := router.Group("/api", middleware.AuthRequired())
	protectedApiRouter.GET("/john", func(c *gin.Context) {
		user, _ := c.Get("User")
		c.JSON(200, "Authenticated as: "+user.(string))
	})
	protectedApiRouter.PUT("/meal", func(c *gin.Context) {
		var meal model.Meal
		err := c.MustBindWith(&meal, binding.JSON)
		if err != nil {
			fmt.Println("Could not parse meal", err)
			c.JSON(400, err.Error())
			return
		}

		err = validateMeal(meal, true)
		if err != nil {
			fmt.Println("Error:", err)
			c.JSON(400, "invalid meal object")
			return
		}

		err = mealDao.Update(meal)
		if err != nil {
			fmt.Println("Error when updating meal:", err)
			c.JSON(500, "error when updating meal: "+err.Error())
			return
		}

		c.JSON(200, "updated")
	})
	protectedApiRouter.POST("/meal", func(c *gin.Context) {
		var meal model.Meal
		err := c.MustBindWith(&meal, binding.JSON)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = validateMeal(meal, false)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		mealDao.Insert(&meal)

		c.JSON(201, "created")
	})
	protectedApiRouter.GET("/meal/:id", func(c *gin.Context) {
		idAsString := c.Param("id")
		id, err := strconv.ParseInt(idAsString, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			c.JSON(400, "ID can't be parsed as int")
			return
		}

		meal := mealDao.GetById(id)

		if meal != nil {
			c.JSON(200, *meal)
		} else {
			c.JSON(404, "not found")
		}
	})
	protectedApiRouter.DELETE("/meal/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			c.JSON(400, "id missing")
			return
		}

		mealDeleted, err := mealDao.Delete(id)
		if err != nil {
			fmt.Println("Error deleting meal:", err)
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
