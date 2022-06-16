package routes

import (
	"yemeksepeti/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(a *fiber.App) {
	//publics
	a.Post("/api/register", controllers.Register)
	a.Post("/api/login", controllers.Login)
	a.Get("/api/logout", controllers.Logout)
	//privates
	a.Post("/api/order", controllers.CreateOrder)
	a.Get("/api/order", controllers.ListOrders)
	a.Get("/api/order/:data", controllers.ListOrder)
}
