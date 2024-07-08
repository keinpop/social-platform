package auth

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"mai-platform/internal/clients/db"
	"net/http"
	"net/mail"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	HashCost    = 14
	identityKey = "id"
)

type Auth struct {
	config *Config
	DB     *db.DB
}

func NewAuth(cfg *Config) *Auth {
	return &Auth{
		config: cfg,
		DB:     db.NewDB(&cfg.DB),
	}
}

func (auth *Auth) Init() error {
	if err := auth.DB.Init(); err != nil {
		return err
	}

	return nil
}

func (auth *Auth) hashPassword(login, password string) (string, error) {
	// s := sha256.Sum256([]byte(login + password))
	// TODO: нормально сделать
	return login + password, nil
}

type LoginPassword struct {
	Login    string `form:"login" json:"login" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (auth *Auth) Register(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[error] Failed to read body: %v", err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	var lp LoginPassword
	err = json.Unmarshal(jsonData, &lp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	_, err = mail.ParseAddress(lp.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad email",
		})
		return
	}

	hp, err := auth.hashPassword(lp.Login, lp.Password)
	if err != nil {
		log.Printf("[error] Failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	err = auth.DB.Register(lp.Login, hp)
	switch {
	case err == nil:
		c.JSON(http.StatusCreated, "")
	case errors.Is(err, gorm.ErrCheckConstraintViolated) || errors.Is(err, gorm.ErrDuplicatedKey):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
	default:
		log.Printf("Failed to register user: %v", err)
		c.JSON(http.StatusInternalServerError, "")
	}
}

// CheckToken should be called after request is authenticated.
func (auth *Auth) CheckToken(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"login": claims[identityKey],
	})
}

func (auth *Auth) getAuthentificator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var lp LoginPassword
		log.Print("HUY")
		if err := c.ShouldBind(&lp); err != nil {
			log.Print(err)
			return "", jwt.ErrMissingLoginValues
		}

		hp, err := auth.hashPassword(lp.Login, lp.Password)
		if err != nil {
			log.Printf("[error] Failed to hash password: %v", err)
			c.JSON(http.StatusInternalServerError, "")
		}

		ok, err := auth.DB.CheckHash(lp.Login, hp)
		log.Print(ok, err)
		if ok && err == nil {
			return lp.Login, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}

func (auth *Auth) GetJWTMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(auth.config.SecretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{
				identityKey: data,
			}
		},

		Authenticator: auth.getAuthentificator(),
		Authorizator:  func(data interface{}, c *gin.Context) bool { return true },
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}
