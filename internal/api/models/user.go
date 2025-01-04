package models

const (
  UserStatusConnected = iota
  UserStatusInRoom = iota
)

type UserModel struct {
  Name string `json:"name"`
  Status int `json:"status"`
  Score int `json:"score"`
}

func MakeUserModel(name string) *UserModel {
  return &UserModel{
    Name: name,
    Status: UserStatusConnected,
    Score: 0,
  }
}
