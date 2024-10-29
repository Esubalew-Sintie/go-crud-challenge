// internal/adapters/postgres/repository.go
package gormdb

import (
	"errors"
	"go-crud-challenge/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresPersonRepository struct {
	db *gorm.DB
}

func NewPostgresPersonRepository(db *gorm.DB) *PostgresPersonRepository {
	db.AutoMigrate(&domain.Person{})
	return &PostgresPersonRepository{db: db}
}

func (r *PostgresPersonRepository) GetAll() ([]domain.Person, error) {
	var persons []domain.Person
	if err := r.db.Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

func (r *PostgresPersonRepository) GetByID(id string) (domain.Person, error) {
	var person domain.Person
	if err := r.db.First(&person, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Person{}, errors.New("person not found")
		}
		return domain.Person{}, err
	}
	return person, nil
}

func (r *PostgresPersonRepository) Create(person domain.Person) (domain.Person, error) {
	person.ID = uuid.NewString()
	if err := r.db.Create(&person).Error; err != nil {
		return domain.Person{}, err
	}
	return person, nil
}

func (r *PostgresPersonRepository) Update(id string, person domain.Person) (domain.Person, error) {
	person.ID = id
	if err := r.db.Save(&person).Error; err != nil {
		return domain.Person{}, err
	}
	return person, nil
}

func (r *PostgresPersonRepository) Delete(id string) error {
	if err := r.db.Delete(&domain.Person{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
