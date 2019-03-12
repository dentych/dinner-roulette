package database

import (
	"database/sql"
	"github.com/dentych/dinner-dash/models"
)

type UserDao struct {
}

func (dao *UserDao) Insert(user models.User) (int, error) {
	db := GetConnection()

	query := "INSERT INTO public.user (email, passwordHash) VALUES ($1, $2) RETURNING id"
	result := db.QueryRowx(query, user.Email, user.PasswordHash)
	var value int
	err := result.Scan(&value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (dao *UserDao) GetUserByEmail(email string) (models.User, error) {
	db := GetConnection()

	var user models.User
	query := "SELECT id, email, passwordHash FROM public.user WHERE email = $1"
	err := db.Get(&user, query, email)
	if err != nil {
		return user, err
	}

	return user, nil
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

	query := "INSERT INTO public.session (userId, sessionId, validTo) VALUES ($1, $2, NOW() + INTERVAL '1 DAY')"
	_, err := db.Exec(query, userId, session)

	return err
}

func (dao *UserDao) CheckSession(userId int, sessionId string) error {
	db := GetConnection()

	var session models.Session
	query := "SELECT id, userId, sessionId, validTo FROM public.session WHERE userId = $1 AND sessionId = $2 AND validTo > NOW()"
	err := db.Get(&session, query, userId, sessionId)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) DeleteSessionsByUserId(userId int) error {
	db := GetConnection()

	query := "DELETE FROM public.session WHERE userId = $1"
	_, err := db.Exec(query, userId)

	return err
}