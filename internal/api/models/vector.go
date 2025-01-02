package models


type Vector struct {
  X int
  Y int
}

func (v *Vector) Add(other *Vector) {
  v.X += other.X
  v.Y += other.Y
}

func (v *Vector) Multiply(multiplicator int) {
  v.X *= multiplicator
  v.Y *= multiplicator
}
