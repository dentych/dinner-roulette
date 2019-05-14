package main

import (
	"fmt"
	"github.com/dentych/dinner-dash/config"
	"github.com/dentych/dinner-dash/database"
	//"github.com/dentych/dinner-dash/models"
)

func main() {
	configuration := config.FromEnv()
	database.Init(configuration.DbConfig)
//	ingredientDato := database.IngredientDao{}
//	recipeDao := database.RecipeDao{}
	recipeIngredientDao := database.RecipeIngredientDao{}

	data, _ := recipeIngredientDao.GetByRecipe(3)
	fmt.Printf("ID: %d\n", data[0].ID)
	fmt.Printf("Amount: %d\n", data[0].Amount)
	fmt.Printf("Unit: %s\n", data[0].Unit)
	fmt.Printf("Name: %s\n", data[0].Name)
	fmt.Printf("KCAL: %d\n", data[0].KCAL)

/*

	first_ing := ingredientDato.GetById(1)
	first_rec := recipeDao.GetById(1,3)

	fmt.Println(first_ing)
	fmt.Println(first_rec)

	var recIng models.RecipeIngredient
	recIng.Amount = 2
	recIng.Unit = "Kg"

	result := recipeIngredientDao.Insert(&recIng, first_rec.ID, first_ing.ID)

	fmt.Println(result)*/
/*	var ingredient models.Ingredient
	ingredient.Name = "Apple"
	ingredient.KCAL = 20

	var id = ingredientDato.Insert(&ingredient)
	fmt.Println(id)


	ingredients, err := ingredientDato.GetLikePattern("App")
	if err != nil{
		fmt.Println(err)
	}

	for index,element := range ingredients {
          "github.com/dentych/dinner-dash/logging"
fmt.Println(index)
        fmt.Println(element)
  }

	fmt.Println("Getting id 1")
	ingredient2 := ingredientDato.GetById(1)
	fmt.Println(ingredient2)

	ingredient2.KCAL = 2002
	errz := ingredientDato.Update(*ingredient2)

	if errz != nil{
		fmt.Println("Failed to update.")
	}
	ingredientDato.Delete(2)

*/
	return
}
