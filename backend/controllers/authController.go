package controllers

import (
	"database/sql"
	"github.com/dentych/dinner-dash/config"
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
			logging.Error.Println("Database error: ", err)
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

	ctl.setUserSessionCookies(ctx, user.ID, sess)
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

	if isEmptyOrWhitespace(u.FirstName) || isEmptyOrWhitespace(u.LastName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "first and last name should be filled out"})
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

	userToSave := models.User{Email: u.Email, PasswordHash: hash, FirstName: u.FirstName, LastName: u.LastName}

	userId, err := ctl.userDao.Insert(userToSave)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save user"})
		return
	}

	session, err := security.GenerateSession()
	if err == nil {
		err = ctl.userDao.InsertSession(userId, session)
		if err == nil {
			ctl.setUserSessionCookies(ctx, userId, session)
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

func (ctl *AuthController) Logout(ctx *gin.Context) {
	ctl.unsetUserSessionCookies(ctx)

	ctx.Status(http.StatusOK)
}

func (ctl *AuthController) setUserSessionCookies(ctx *gin.Context, userId int, session string) {
	setCookie(ctx, "userId", strconv.Itoa(userId), "/api/token", ctl.cookieHost, 86400)
	setCookie(ctx, "session", session, "/api/token", ctl.cookieHost, 86400)
}

func (ctl *AuthController) unsetUserSessionCookies(ctx *gin.Context) {
	setCookie(ctx, "userId", "", "/api/token", ctl.cookieHost, -1)
	setCookie(ctx, "session", "", "/api/token", ctl.cookieHost, -1)
}

func setCookie(ctx *gin.Context, name, value, path, host string, maxAge int) {
	ctx.SetCookie(name, value, maxAge, path, host, config.IsProd(), true)
}

func hashFromPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(hash), err
}

func isEmptyOrWhitespace(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

type user struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
