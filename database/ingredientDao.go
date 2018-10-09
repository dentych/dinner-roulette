package database

import (
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/model"
)

type IngredientDao struct {
}

func (dao *IngredientDao) Insert(ingredient *model.Ingredient) error {
	db, err := getConnection()
	if err != nil {
		logging.Error.Printf("Error getting database connection: %s", err)
		return err
	}

	sql := "INSERT INTO ingredient (name, unit) VALUES ($1, $2) RETURNING id"
	if err = db.Get(&ingredient.Id, sql, ingredient.Name, ingredient.Unit); err != nil {
		logging.Error.Printf("Error inserting ingredient: %s", err)
		return err
	}

	return nil
}

func (dao *IngredientDao) Update(ingredient model.Ingredient) error {
	db, err := getConnection()
	if err != nil {
		logging.Error.Printf("Error: %s", err)
		return err
	}

	sql := "UPDATE ingredient SET name = $1, unit = $2"
	if _, err := db.Exec(sql, ingredient.Name, ingredient.Unit); err != nil {
		logging.Error.Printf("Error: %s", err)
		return err
	}

	return nil
}
func (dao *IngredientDao) GetAll() ([]model.Ingredient, error) {
	db, err := getConnection()
	if err != nil {
		logging.Error.Printf("Error: %s", err)
		return nil, err
	}

	ingredients := make([]model.Ingredient, 0, 15)
	sql := "SELECT id, name, unit FROM ingredient"
	if err = db.Select(&ingredients, sql); err != nil {
		logging.Error.Printf("Error: %s", err)
		return nil, err
	}

	return ingredients, nil
}