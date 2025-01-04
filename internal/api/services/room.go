package services

import (
  "log"

  "github.com/romakot321/game-backend/internal/api/models"
)

type RoomService interface {
  Authenticate(msg models.MessageAuthenticateData) *models.Room
  List() []*models.Room
}

type roomService struct {
  rooms map[string]*models.Room
}

func (s *roomService) Authenticate(msg models.MessageAuthenticateData) *models.Room {
  room, ok := s.rooms[msg.Room]
  if ok {
    return room
  }
  users := make([]*models.UserModel, 0)
  room = &models.Room{Name: msg.Room, Users: users}
  s.rooms[msg.Room] = room
  log.Println("Create room " + msg.Room)
  return room
}

func (s *roomService) List() []*models.Room {
  resp := make([]*models.Room, len(s.rooms))
  for _, room := range s.rooms {
    resp = append(resp, room)
  }
  return resp
}

func NewRoomService() RoomService {
  rooms := make(map[string]*models.Room, 0)
  return &roomService{rooms: rooms}
}
