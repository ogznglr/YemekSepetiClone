package controllers

import (
	"fmt"
	"strconv"
	"time"
	"yemeksepeti/database"
	"yemeksepeti/models"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.JSON(map[string]string{
			"hata var": "yanlış veri geldi",
		})
	}

	claims, err := UserValidation(c)
	if err != nil {
		return c.JSON(map[string]string{
			"sıkıntı var": fmt.Sprintf("%d", fiber.StatusBadRequest),
		})
	}
	var order models.Order
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if strconv.Itoa(int(user.Id)) != claims.Issuer {
		return c.JSON(map[string]string{
			"hata var": "kullanıcı bulunamadı",
		})
	}
	order.Userid = claims.Issuer
	order.Restoranid = data["restoranid"]
	order.Productid = data["productid"]
	order.ExpTime = time.Now().Add(time.Hour).String()
	database.DB.Create(&order)
	return c.JSON(map[string]interface{}{
		"ordder has been saved successfully": &order,
	})
}
func ListOrders(c *fiber.Ctx) error {
	claims, err := UserValidation(c)
	if err != nil {
		return c.JSON(map[string]string{
			"hata mevcut": "kullanıcı bulunamadı",
		})
	}

	var orders []models.Order

	database.DB.Where("userid = ?", claims.Issuer).Find(&orders)
	return c.JSON(&orders)

}

func ListOrder(c *fiber.Ctx) error {
	data := c.Params("data")
	var order models.Order
	database.DB.Where("id = ?", data).Find(&order)
	return c.JSON(order)
}
