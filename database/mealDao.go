package database

import (
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/model"
	_ "github.com/lib/pq"
)

type MealDao struct {
}

func (dao *MealDao) Insert(m *model.Meal) int {
	db, err := GetConnection()
	if err != nil {
		return -1
	}

	result := db.MustExec("INSERT INTO meal (name, url) VALUES ($1, $2)", m.Name, m.Url)
	rows, err := result.RowsAffected()
	if err != nil {
		return -1
	}

	return int(rows)
}

func (dao *MealDao) GetAll() []model.Meal {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Printf("Could not connect to database: %s", err)
		return make([]model.Meal, 0)
	}

	var meals []model.Meal
	err = db.Select(&meals, "SELECT * FROM meal")

	if err != nil {
		logging.Error.Println(err)
		return make([]model.Meal, 0)
	}

	return meals
}

func (dao *MealDao) GetById(id int64) *model.Meal {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Println(err)
		return nil
	}

	var meal model.Meal
	err = db.Get(&meal, "SELECT * FROM meal WHERE id = $1", id)
	if err != nil {
		logging.Error.Println(err)
		return nil
	}

	return &meal
}

func (dao *MealDao) Update(meal model.Meal) error {
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

func (dao *MealDao) Delete(id int64) (bool, error) {
	db, err := GetConnection()
	if err != nil {
		logging.Error.Println(err)
		return false, err
	}

	sql := "DELETE FROM meal WHERE id = $1"
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
