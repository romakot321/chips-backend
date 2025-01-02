package repositories


import (
  "github.com/romakot321/game-backend/internal/api/models"
)

type EntityRepository interface {
  Add(entity *models.EntityModel)
  Update(schema models.EntityModel) *models.EntityModel
}

type entityRepository struct {
  entities []*models.EntityModel
}

func (r *entityRepository) Add(entity *models.EntityModel) {
  r.entities = append(r.entities, entity)
}

func (r entityRepository) Update(schema models.EntityModel) *models.EntityModel {
  return nil
  for _, entity := range r.entities {
    if schema.Position != nil {
      entity.Position.Add(schema.Position)
    }
    return entity
  }
  return nil
}

func NewEntityRepository() EntityRepository {
  entities := make([]*models.EntityModel, 0)
  return &entityRepository{entities: entities}
}
