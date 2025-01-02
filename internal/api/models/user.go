package models

const (
  UserStatusConnected = iota
  UserStatusInRoom = iota
)

type UserModel struct {
  Name string `json:"name"`
  Status int `json:"status"`
  Position *Vector `json:"position"`
}

func MakeUserModel(name string) *UserModel {
  return &UserModel{
    Name: name,
    Status: UserStatusConnected,
    Position: &Vector{X: 0, Y: 0},
  }
}
