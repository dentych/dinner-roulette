package database

import (
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/model"
	_ "github.com/lib/pq"
)

type MealDao struct {
}

func (dao *MealDao) Insert(username string, m *model.Meal) int {
	db, err := GetConnection()
	if err != nil {
		return -1
	}

	sql := `INSERT INTO public.meal (name, url, userid)
VALUES ($1, $2, (SELECT id FROM public.user WHERE username=$3))`
	result := db.MustExec(sql, m.Name, m.Url, username)
	rows, err := result.RowsAffected()
	if err != nil {
		return -1
	}

	return int(rows)
}

func (dao *MealDao) GetAll(username string) []model.Meal {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Printf("Could not connect to database: %s", err)
		return make([]model.Meal, 0)
	}

	var meals []model.Meal
	sql := "SELECT id, name, url FROM meal WHERE userid = (SELECT id FROM public.user WHERE username = $1)"
	err = db.Select(&meals, sql, username)

	if err != nil {
		logging.Error.Println(err)
		return make([]model.Meal, 0)
	}

	return meals
}

func (dao *MealDao) GetById(username string, id int64) *model.Meal {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Println(err)
		return nil
	}

	var meal model.Meal
	sql := "SELECT id, name, url FROM meal WHERE id = $1 AND userid = (SELECT id FROM public.user WHERE username = $2)"
	err = db.Get(&meal, sql, id, username)
	if err != nil {
		logging.Error.Println(err)
		return nil
	}

	return &meal
}

func (dao *MealDao) Update(username string, meal model.Meal) error {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Println(err)
		return err
	}

	sql := "UPDATE meal SET name = $1, url = $2 WHERE id = $3"
	_, err = db.Exec(sql, meal.Name, meal.Url, meal.ID)
	if err != nil {
		logging.Error.Println(err)
		return err
	}

	return nil
}

func (dao *MealDao) Delete(username string, id int64) (bool, error) {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Println(err)
		return false, err
	}

	sql := "DELETE FROM meal WHERE id = $1 AND userid = (SELECT id FROM public.user WHERE username = $2)"
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
