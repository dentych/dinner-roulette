package handlers

import (
	"fmt"
	"github.com/dentych/dinner-dash/database"
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

type RecipeHandler struct {
	RecipeDao database.RecipeDao
}

func NewRecipeHandler(recipeDao database.RecipeDao) *RecipeHandler {
	controller := RecipeHandler{RecipeDao: recipeDao}
	return &controller
}

func (rc *RecipeHandler) MealPlan(days int) {

}

func (rc *RecipeHandler) CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
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

	rc.RecipeDao.Insert(c.GetInt("uid"), &recipe)

	c.JSON(201, recipe)
}

func (rc RecipeHandler) GetRecipes(c *gin.Context) {
	recipes, err := rc.RecipeDao.GetAll(c.GetInt("uid"))
	if err != nil {
		c.JSON(500, "error while getting recipes")
		return
	}

	c.JSON(200, recipes)
}

func (rc RecipeHandler) GetRecipeById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logging.Error.Printf("Error: %s", err)
		c.JSON(400, "ID can't be parsed as int")
		return
	}

	recipe := rc.RecipeDao.GetById(int64(c.GetInt("uid")), id)

	if recipe != nil {
		c.JSON(200, *recipe)
	} else {
		c.JSON(404, "not found")
	}
}

func (rc RecipeHandler) UpdateRecipe(c *gin.Context) {
	uid := c.GetInt("uid")
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid ID"})
		return
	}

	var recipe models.Recipe
	err = c.MustBindWith(&recipe, binding.JSON)
	if err != nil {
		logging.Error.Printf("Could not parse recipe: %s", err)
		c.JSON(400, err.Error())
		return
	}

	err = validateRecipe(recipe, false)
	if err != nil {
		logging.Info.Printf("Could not validate recipe object: %s", err)
		c.JSON(400, "invalid recipe object")
		return
	}

	recipe.ID = id
	err = rc.RecipeDao.Update(uid, recipe)
	if err != nil {
		logging.Error.Printf("Error when updating recipe: %s", err)
		c.JSON(500, "error when updating recipe: "+err.Error())
		return
	}

	c.JSON(200, "updated")
}

func (rc RecipeHandler) DeleteRecipe(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logging.Error.Printf("Error: %s", err)
		c.JSON(400, "invalid or missing id missing from URL path (/api/recipe/{id})")
		return
	}

	recipeDeleted, err := rc.RecipeDao.Delete(c.GetInt("uid"), id)
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
}

func validateRecipe(recipe models.Recipe, checkId bool) error {
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
