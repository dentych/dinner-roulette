package controllers

import (
	"github.com/dentych/dinner-dash/database"
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/models"
	"github.com/dentych/dinner-dash/security"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type AuthController struct {
	userDao database.UserDao
}

func NewAuthController(userDao database.UserDao) *AuthController {
	return &AuthController{userDao: userDao}
}

func (ctl *AuthController) Login(ctx *gin.Context) {
	u := user{}
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse the user sent by client"})
		logging.Error.Printf("Unable to parse the information sent by client")
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "nope"})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error.", "error_code": "A1"})
		logging.Error.Printf("Error: %v", err)
		return
	}
	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email exists"})
		return
	}

	salt := getSalt()
	hash, err := hashFromPassword(u.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred.", "error_code": "A2"})
		return
	}

	userToSave := models.User{Email: u.Email, PasswordHash: hash, Salt: salt}

	userId, err := ctl.userDao.Insert(userToSave)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save user"})
		return
	}

	session, err := security.GenerateSession()
	if err == nil {
		err = ctl.userDao.InsertSession(userId, session)
		if err == nil {
			ctx.SetCookie("userid", string(userId), 0, "/", "localhost", false, true)
			ctx.SetCookie("session", session, 0, "/", "localhost", false, true)
		}
	}

	ctx.Status(http.StatusOK)
}

func hashFromPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(hash), err
}

func getSalt() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(100000000)
}

func isEmptyOrWhitespace(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

type user struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
