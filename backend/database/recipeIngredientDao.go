package database

import (
  "fmt"
	"github.com/dentych/dinner-dash/models"
//  "github.com/dentych/dinner-dash/logging"
	_ "github.com/lib/pq"
)
type RecipeIngredientDao struct {
}


func (dao *RecipeIngredientDao) InsertIngredientMapping(amount int, unit string, recipe_id int, ingredient_id int) int {
	db := GetConnection()

	sql := `INSERT INTO public.recipe_ingredient (ingredient_id, recipe_id, amount, unit)
VALUES ($1, $2, $3, $4)`
	result := db.QueryRowx(sql, ingredient_id, recipe_id, amount, unit)
	var value int
	err := result.Scan(&value)
	if err != nil {
    fmt.Println(err)
		return -1
	}

	return 1
}

func (dao *RecipeIngredientDao) GetByRecipe(rid int) ([]models.RecipeIngredient, error) {
  db := GetConnection()
  var recing = make([]models.RecipeIngredient, 0, 0)
  sql := "SELECT recipe_ingredient.id, recipe_ingredient.amount, recipe_ingredient.unit, ingredient.id, ingredient.name, ingredient.kcal FROM public.recipe_ingredient JOIN ingredient ON ingredient_id = ingredient.id WHERE recipe_id = $1"
  err := db.Select(&recing, sql, rid)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return recing, nil
}

func (dao *RecipeIngredientDao) GetAll() ([]models.RecipeIngredient, error) {
	db := GetConnection()

	var recing = make([]models.RecipeIngredient, 0, 0)
	sql := "SELECT * FROM public.recipe_ingredient ORDER BY id"
	err := db.Select(&recing, sql)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return recing, nil
}

func (dao *RecipeIngredientDao) GetById(id int64) *models.RecipeIngredient {
	db := GetConnection()

	var recing models.RecipeIngredient

	sql := "SELECT * FROM public.recipe_ingredient WHERE id = $1"
	err := db.Get(&recing, sql, id)
	if err != nil {
    fmt.Println(err)
		return nil
	}

	return &recing
}
/*
func (dao *RecipeIngredientDao) Update(ingredient models.RecipeIngredient) error {
	db := GetConnection()

	sql := "UPDATE public.ingredient SET name = $1, kcal = $2 WHERE id = $3"

	_, err := db.Exec(sql, ingredient.Name, ingredient.KCAL, ingredient.ID)
	if err != nil {
		logging.Error.Println(err)
		return err
	}

	return nil
}

func (dao *RecipeIngredientDao) Delete(id int64) (bool, error) {
	db := GetConnection()

	sql := "DELETE FROM public.ingredient WHERE id = $1"
	result, err := db.Exec(sql, id)
	if err != nil {
		logging.Error.Println(err)
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logging.Error.Println(err)
		return false, err
	}
	if rowsAffected < 1 {
		return false, nil
	}
	return true, nil
}
*/
