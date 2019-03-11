package database

import (
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/models"
	_ "github.com/lib/pq"
)

type RecipeDao struct {
}

func (dao *RecipeDao) Insert(username string, m *models.Recipe) int {
	db := GetConnection()

	sql := `INSERT INTO public.recipe (name, url, userid)
VALUES ($1, $2, (SELECT id FROM public.user WHERE username=$3)) RETURNING id`
	result := db.QueryRowx(sql, m.Name, m.Url, username)
	var value int
	err := result.Scan(&value)
	if err != nil {
		return -1
	}
	m.ID = value

	return 1
}

func (dao *RecipeDao) GetAll(username string) ([]models.Recipe, error) {
	db := GetConnection()

	var recipes = make([]models.Recipe, 0, 0)
	sql := "SELECT id, name, url FROM recipe WHERE userid = (SELECT id FROM public.user WHERE username = $1)"
	err := db.Select(&recipes, sql, username)

	if err != nil {
		logging.Error.Println(err)
		return nil, err
	}

	return recipes, nil
}

func (dao *RecipeDao) GetById(username string, id int64) *models.Recipe {
	db := GetConnection()

	var recipe models.Recipe
	sql := "SELECT id, name, url FROM recipe WHERE id = $1 AND userid = (SELECT id FROM public.user WHERE username = $2)"
	err := db.Get(&recipe, sql, id, username)
	if err != nil {
		logging.Error.Println(err)
		return nil
	}

	return &recipe
}

func (dao *RecipeDao) Update(username string, recipe models.Recipe) error {
	db := GetConnection()

	sql := "UPDATE recipe SET name = $1, url = $2 WHERE id = $3"
	_, err := db.Exec(sql, recipe.Name, recipe.Url, recipe.ID)
	if err != nil {
		logging.Error.Println(err)
		return err
	}

	return nil
}

func (dao *RecipeDao) Delete(username string, id int64) (bool, error) {
	db := GetConnection()

	sql := "DELETE FROM recipe WHERE id = $1 AND userid = (SELECT id FROM public.user WHERE username = $2)"
	result, err := db.Exec(sql, id, username)
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
