package service

import (
	"crypto/sha256"
	"encoding/hex"
	"example/models"
	"example/repository"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var APPLICATION_NAME = "Leave Application"
var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("bematrix1")

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
}

func Login(c echo.Context) error {
	U := new(models.LoginModel)
	if err := c.Bind(U); err != nil {
		return nil
	}
	result := repository.ReadUserByUsername(U.Username)

	h := sha256.New()
	h.Write([]byte(U.Password))
	passwordHash := hex.EncodeToString(h.Sum(nil))

	if result != nil {
		for _, each := range result {
			isValid := each.Password == passwordHash

			if isValid {
				claims := MyClaims{
					StandardClaims: jwt.StandardClaims{
						Issuer:    APPLICATION_NAME,
						ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
					},
					Username: U.Username,
				}

				token := jwt.NewWithClaims(
					JWT_SIGNING_METHOD,
					claims,
				)

				signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
				if err != nil {
					return c.JSON(http.StatusForbidden, err)
				}

				result := models.ResponseLogin{
					User:    each,
					Token:   signedToken,
					Message: "Login Success",
				}

				return c.JSON(http.StatusOK, result)

			} else {
				response := models.ResponseLogin{
					Token:   "",
					Message: "Wrong Password!",
				}
				return c.JSON(http.StatusForbidden, response)
			}
		}
	} else {
		response := models.ResponseLogin{
			Token:   "",
			Message: "Username does not registered!",
		}
		return c.JSON(http.StatusForbidden, response)
	}

	return nil
}
