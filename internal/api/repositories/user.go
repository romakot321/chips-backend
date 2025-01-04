package repositories


import (
  "github.com/romakot321/game-backend/internal/api/models"
)

type UserRepository interface {
  Add(user *models.UserModel)
  Get(name string) *models.UserModel
  Update(schema models.UserModel) *models.UserModel
  GetList() []*models.UserModel
}

type userRepository struct {
  users []*models.UserModel
}

func (r *userRepository) Add(user *models.UserModel) {
  r.users = append(r.users, user)
}

func (r *userRepository) Get(name string) *models.UserModel {
  for _, user := range r.users {
    if user.Name == name {
      return user
    }
  }
  return nil
}

func (r userRepository) Update(schema models.UserModel) *models.UserModel {
  for _, user := range r.users {
    if user.Name != schema.Name {
      continue
    }
    if schema.Score != 0 {
      user.Score += schema.Score
    }
    return user
  }
  return nil
}

func (r userRepository) GetList() []*models.UserModel {
  return r.users
}

func NewUserRepository() UserRepository {
  users := make([]*models.UserModel, 0)
  return &userRepository{users: users}
}
