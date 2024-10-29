// internal/service/person_service.go
package service

import (
	"go-crud-challenge/internal/domain"
	"go-crud-challenge/internal/ports"
)

type PersonService struct {
	repo ports.PersonRepository
}

func NewPersonService(repo ports.PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) GetAllPersons() ([]domain.Person, error) {
	return s.repo.GetAll()
}

func (s *PersonService) GetPersonByID(id string) (domain.Person, error) {
	return s.repo.GetByID(id)
}

func (s *PersonService) CreatePerson(person domain.Person) (domain.Person, error) {
	return s.repo.Create(person)
}

func (s *PersonService) UpdatePerson(id string, person domain.Person) (domain.Person, error) {
	return s.repo.Update(id, person)
}

func (s *PersonService) DeletePerson(id string) error {
	return s.repo.Delete(id)
}
