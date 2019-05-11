package database

import (
  "fmt"
	"github.com/dentych/dinner-dash/models"
  "github.com/dentych/dinner-dash/logging"
	_ "github.com/lib/pq"
)
type IngredientDao struct {
}

func (dao *IngredientDao) Insert(m *models.Ingredient) int {
	db := GetConnection()

	sql := `INSERT INTO public.ingredient (name, kcal)
VALUES ($1, $2) RETURNING id`
	result := db.QueryRowx(sql, m.Name, m.KCAL)
	var value int
	err := result.Scan(&value)
	if err != nil {
    fmt.Println(err)
		return -1
	}
	m.ID = value

	return 1
}

func (dao *IngredientDao) GetAll() ([]models.Ingredient, error) {
	db := GetConnection()

	var ingredients = make([]models.Ingredient, 0, 0)
	sql := "SELECT id, name, kcal FROM public.ingredient ORDER BY name"
	err := db.Select(&ingredients, sql)

	if err != nil {
		logging.Error.Println(err)
		return nil, err
	}

	return ingredients, nil
}

func (dao *IngredientDao) GetLikePattern(pattern string) ([]models.Ingredient, error) {
	db := GetConnection()
  wildstring := pattern + "%"
	var ingredients = make([]models.Ingredient, 0, 0)
	sql := "SELECT * FROM public.ingredient WHERE name LIKE $1"
	err := db.Select(&ingredients, sql,wildstring)

	if err != nil {
		logging.Error.Println(err)
		return nil, err
	}

	return ingredients, nil
}

func (dao *IngredientDao) GetById(id int64) *models.Ingredient {
	db := GetConnection()

	var ingredient models.Ingredient

	sql := "SELECT * FROM public.ingredient WHERE id = $1"
	err := db.Get(&ingredient, sql, id)
	if err != nil {
    fmt.Println(err)
    //logging.Error.Println(err)
		return nil
	}


	return &ingredient
}

func (dao *IngredientDao) Update(ingredient models.Ingredient) error {
	db := GetConnection()

	sql := "UPDATE public.ingredient SET name = $1, kcal = $2 WHERE id = $3"

	_, err := db.Exec(sql, ingredient.Name, ingredient.KCAL, ingredient.ID)
	if err != nil {
		logging.Error.Println(err)
		return err
	}

	return nil
}

func (dao *IngredientDao) Delete(id int64) (bool, error) {
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
