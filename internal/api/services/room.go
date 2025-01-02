package services

import (
  "github.com/romakot321/game-backend/internal/api/repositories"
  "github.com/romakot321/game-backend/internal/api/models"
)

type Room struct {
  entityRepository repositories.EntityRepository
  users []*models.UserModel
  Name string
}

type RoomService interface {
  Authenticate(msg models.MessageAuthenticateData) *Room
  GetList() []*Room
}

type roomService struct {
  rooms map[string]*Room
}

func (s *roomService) Authenticate(msg models.MessageAuthenticateData) *Room {
  room, ok := s.rooms[msg.Room]
  if ok {
    return room
  }
  entityRepository := repositories.NewEntityRepository()
  users := make([]*models.UserModel, 0)
  room = &Room{entityRepository: entityRepository, Name: msg.Room, users: users}
  s.rooms[msg.Room] = room
  return room
}

func (s *roomService) GetList() []*Room {
  resp := make([]*Room, len(s.rooms))
  for _, room := range s.rooms {
    resp = append(resp, room)
  }
  return resp
}

func NewRoomService() RoomService {
  rooms := make(map[string]*Room)
  return &roomService{rooms: rooms}
}

func (r *Room) AddEntity(model *models.EntityModel) {
  r.entityRepository.Add(model)
}

func (r *Room) AddUser(model *models.UserModel) {
  r.users = append(r.users, model)
}
