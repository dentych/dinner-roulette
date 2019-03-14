package controllers

import (
	"database/sql"
	"github.com/dentych/dinner-dash/database"
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/models"
	"github.com/dentych/dinner-dash/security"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
)

type AuthController struct {
	userDao    database.UserDao
	cookieHost string
}

func NewAuthController(userDao database.UserDao, cookieHost string) *AuthController {
	return &AuthController{userDao: userDao, cookieHost: cookieHost}
}

func (ctl *AuthController) Login(ctx *gin.Context) {
	u := user{}
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse the user sent by client"})
		logging.Error.Printf("Unable to parse the information sent by client")
		return
	}

	if isEmptyOrWhitespace(u.Email) || isEmptyOrWhitespace(u.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email and password have to be filled out"})
		return
	}

	user, err := ctl.userDao.GetUserByEmail(u.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong email or password"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "database error", "error_code": "A1"})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(u.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong email or password"})
		return
	}

	sess, err := security.GenerateSession()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error when creating session for user", "error_code": "A4"})
		return
	}

	err = ctl.userDao.InsertSession(user.ID, sess)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error when creating session for user", "error_code": "A4"})
		return
	}

	setUserSessionCookie(ctx, user.ID, sess, ctl.cookieHost)
	ctx.Status(http.StatusOK)
}

func (ctl *AuthController) Register(ctx *gin.Context) {
	u := user{}
	err := ctx.BindJSON(&u)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse the information sent by client."})
		logging.Error.Printf("Unable to parse the information sent by client.")
		return
	}

	if isEmptyOrWhitespace(u.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email may not be empty"})
		return
	}

	if isEmptyOrWhitespace(u.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "password may not be empty"})
		return
	}

	exists, err := ctl.userDao.EmailExists(u.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error.", "error_code": "A2"})
		logging.Error.Printf("Error: %v", err)
		return
	}
	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email exists"})
		return
	}

	hash, err := hashFromPassword(u.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred.", "error_code": "A3"})
		return
	}

	userToSave := models.User{Email: u.Email, PasswordHash: hash}

	userId, err := ctl.userDao.Insert(userToSave)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save user"})
		return
	}

	session, err := security.GenerateSession()
	if err == nil {
		err = ctl.userDao.InsertSession(userId, session)
		if err == nil {
			setUserSessionCookie(ctx, userId, session, ctl.cookieHost)
		}
	}

	ctx.Status(http.StatusOK)
}

func (ctl *AuthController) Token(ctx *gin.Context) {
	var userIdString, session string
	var err error
	if userIdString, err = ctx.Cookie("userId"); err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}
	if session, err = ctx.Cookie("session"); err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	err = ctl.userDao.CheckSession(userId, session)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	user, err := ctl.userDao.GetUserById(userId)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	token, err := security.CreateJwtAccessToken(userId, user.Email)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access_token": token})
}

func setUserSessionCookie(ctx *gin.Context, userId int, session string, host string) {
	ctx.SetCookie("userId", strconv.Itoa(userId), 0, "/", host, false, true)
	ctx.SetCookie("session", session, 0, "/", host, false, true)
}

func hashFromPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(hash), err
}

func isEmptyOrWhitespace(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

type user struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
