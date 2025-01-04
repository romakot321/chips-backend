package models

type Room struct {
  Users []*UserModel `json:"users"`
  Name string `json:"name"`
  TotalScore int `json:"total_score"`
}

func (r *Room) AddUser(model *UserModel) {
  for _, user := range r.Users {
    if user.Name == model.Name {
      return
    }
  }
  r.Users = append(r.Users, model)
}
