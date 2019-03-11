package database

import (
	"database/sql"
	"github.com/dentych/dinner-dash/models"
)

type UserDao struct {
}

func (dao *UserDao) Insert(user models.User) (int, error) {
	db := GetConnection()

	query := "INSERT INTO public.user (email, passwordHash, salt) VALUES ($1, $2, $3) RETURNING id"
	result := db.QueryRowx(query, user.Email, user.PasswordHash, user.Salt)
	var value int
	err := result.Scan(&value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (dao *UserDao) EmailExists(email string) (bool, error) {
	db := GetConnection()

	query := "SELECT email FROM public.user WHERE email = $1"
	err := db.Get(&email, query, email)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dao *UserDao) InsertSession(userId int, session string) error {
	db := GetConnection()

	query := "INSERT INTO public.sessions (userid, sessionId, validTo) VALUES ($1, $2, NOW() + INTERVAL '1 DAY')"
	_, err := db.Exec(query, userId, session)
	if err != nil {
		return err
	}

	return nil
}