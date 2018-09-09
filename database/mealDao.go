package database

import (
	"fmt"
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
		fmt.Println("Could not connect to database, ", err)
		return make([]model.Meal, 0)
	}

	var meals []model.Meal
	err = db.Select(&meals, "SELECT * FROM meal")

	if err != nil {
		fmt.Println("Error", err)
		return make([]model.Meal, 0)
	}

	return meals
}

func (dao *MealDao) GetById(id int) *model.Meal {
	db, err := GetConnection()
	if err != nil {
		fmt.Println("Error", err)
		return nil
	}

	var meal model.Meal
	err = db.Get(&meal, "SELECT * FROM meal WHERE id = $1", id)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return &meal
}
