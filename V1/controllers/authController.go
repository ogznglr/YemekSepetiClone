package controllers

import (
	"fmt"
	"yemeksepeti/database"
	"yemeksepeti/models"

	"strconv"

	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const (
	key       = 15
	secretKey = "secret"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), key)
	user := models.User{
		Name:     data["name"],
		Surname:  data["surname"],
		Email:    data["email"],
		Password: string(password),
	}
	database.DB.Create(&user)
	return c.JSON(&user)
}

//LOGİN işlemleri
func Login(c *fiber.Ctx) error {
	var data map[string]string //kullanıcının girdiği user bilgileri
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user) //database'den aldığımız user

	if user.Id == 0 {
		return c.SendString(fmt.Sprintf("%d, Not found", fiber.StatusNotFound))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		return c.SendString("Şifre Yanlış")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: jwt.NewTime(1900000000),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return c.SendString("error happened")
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(cookie)
}

func Logout(c *fiber.Ctx) error {
	logoutcookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&logoutcookie)
	return c.JSON(logoutcookie)
}

func UserValidation(c *fiber.Ctx) (*jwt.StandardClaims, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*jwt.StandardClaims), nil
}
