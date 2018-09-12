package database

import (
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/model"
	_ "github.com/lib/pq"
)

type RecipeDao struct {
}

func (dao *RecipeDao) Insert(username string, m *model.Recipe) int {
	db, err := GetConnection()
	if err != nil {
		return -1
	}

	sql := `INSERT INTO public.recipe (name, url, userid)
VALUES ($1, $2, (SELECT id FROM public.user WHERE username=$3))`
	result := db.MustExec(sql, m.Name, m.Url, username)
	rows, err := result.RowsAffected()
	if err != nil {
		return -1
	}

	return int(rows)
}

func (dao *RecipeDao) GetAll(username string) []model.Recipe {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Printf("Could not connect to database: %s", err)
		return make([]model.Recipe, 0)
	}

	var recipes []model.Recipe
	sql := "SELECT id, name, url FROM recipe WHERE userid = (SELECT id FROM public.user WHERE username = $1)"
	err = db.Select(&recipes, sql, username)

	if err != nil {
		logging.Error.Println(err)
		return make([]model.Recipe, 0)
	}

	return recipes
}

func (dao *RecipeDao) GetById(username string, id int64) *model.Recipe {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Println(err)
		return nil
	}

	var recipe model.Recipe
	sql := "SELECT id, name, url FROM recipe WHERE id = $1 AND userid = (SELECT id FROM public.user WHERE username = $2)"
	err = db.Get(&recipe, sql, id, username)
	if err != nil {
		logging.Error.Println(err)
		return nil
	}

	return &recipe
}

func (dao *RecipeDao) Update(username string, recipe model.Recipe) error {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Println(err)
		return err
	}

	sql := "UPDATE recipe SET name = $1, url = $2 WHERE id = $3"
	_, err = db.Exec(sql, recipe.Name, recipe.Url, recipe.ID)
	if err != nil {
		logging.Error.Println(err)
		return err
	}

	return nil
}

func (dao *RecipeDao) Delete(username string, id int64) (bool, error) {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Println(err)
		return false, err
	}

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
