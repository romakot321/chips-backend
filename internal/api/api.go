package api

import (
  "log"

	"github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/contrib/websocket"
  "github.com/romakot321/game-backend/internal/api/repositories"
  "github.com/romakot321/game-backend/internal/api/services"
  "github.com/romakot321/game-backend/internal/api/controllers"
)

func Run() {
  userRepository := repositories.NewUserRepository()
  userService := services.NewUserService(userRepository)
  roomService := services.NewRoomService()
  connectionController := controllers.NewConnectionController(userService, roomService)

  app := fiber.New()
  router := fiber.New()

  app.Mount("/api", router)
  app.Use(func(c *fiber.Ctx) error {
    if websocket.IsWebSocketUpgrade(c) {
      c.Locals("allowed", true)
      return c.Next()
    }
    return fiber.ErrUpgradeRequired
  })
  app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:8000",
    AllowHeaders: "*",
    AllowMethods: "*",
    AllowCredentials: true,
  }))
  router.Route("/game", connectionController.Register)

  log.Fatal(app.Listen(":8000"))
}
