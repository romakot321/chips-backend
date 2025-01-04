package services

import (
  "log"

  "github.com/romakot321/game-backend/internal/api/models"
  "github.com/romakot321/game-backend/internal/api/repositories"
)

type UserService interface {
  Authenticate(msg models.MessageAuthenticateData) *models.UserModel
  AddScore(name string, msg models.MessageChangeScoreData) *models.UserModel
  Win(roomScore int, msg models.MessageWinData) *models.UserModel
  GetList() []*models.UserModel
  ResetCoins() int
  SetUsersCoins(value int)
}

type userService struct {
  userRepository repositories.UserRepository
}

func (s userService) Authenticate(msg models.MessageAuthenticateData) *models.UserModel {
  user := s.userRepository.Get(msg.Username)
  if user == nil {
    user = models.MakeUserModel(msg.Username)
    s.userRepository.Add(user)
    log.Print("Create user with name ", msg.Username)
  }
  return user
}

func (s userService) AddScore(name string, msg models.MessageChangeScoreData) *models.UserModel {
  schema := models.UserModel{
    Name: name,
    Score: msg.Amount,
  }
  return s.userRepository.Update(schema)
}

func (s userService) Win(roomScore int, msg models.MessageWinData) *models.UserModel {
  schema := models.UserModel{
    Name: msg.Name,
    Score: roomScore,
  }
  return s.userRepository.Update(schema)
}

func (s userService) ResetCoins() int {
  total := 0
  for _, user := range s.userRepository.GetList() {
    schema := models.UserModel{
      Name: user.Name,
      Score: -user.Score,
    }
    s.userRepository.Update(schema)
    total += user.Score
  }
  return total
}

func (s userService) SetUsersCoins(value int) {
  for _, user := range s.userRepository.GetList() {
    schema := models.UserModel{
      Name: user.Name,
      Score: value - user.Score,
    }
    s.userRepository.Update(schema)
  }
}

func (s userService) GetList() []*models.UserModel {
  return s.userRepository.GetList()
}

func NewUserService(userRepository repositories.UserRepository) UserService {
  return &userService{userRepository: userRepository}
}
