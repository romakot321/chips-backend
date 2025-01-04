package controllers

import (
  "github.com/gofiber/fiber/v2"
  "github.com/romakot321/game-backend/internal/api/services"
)

type RoomController interface {
  Register(router fiber.Router)
}

type roomController struct {
  roomService services.RoomService
}

func (c *roomController) Register(router fiber.Router) {
  router.Get("/", c.list)
}

func (c *roomController) list(ctx *fiber.Ctx) error {
  rooms := c.roomService.List()
  return ctx.JSON(rooms)
}

func NewRoomController(roomService services.RoomService) RoomController {
  return &roomController{roomService: roomService}
}
