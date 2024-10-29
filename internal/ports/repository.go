// internal/ports/repository.go
package ports

import "go-crud-challenge/internal/domain"

type PersonRepository interface {
	GetAll() ([]domain.Person, error)
	GetByID(id string) (domain.Person, error)
	Create(person domain.Person) (domain.Person, error)
	Update(id string, person domain.Person) (domain.Person, error)
	Delete(id string) error
}
