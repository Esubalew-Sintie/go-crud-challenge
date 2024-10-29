// internal/adapters/memory/repository.go
package memory

import (
	"errors"
	"go-crud-challenge/internal/domain"
	"sync"

	"github.com/google/uuid"
)

type InMemoryPersonRepository struct {
	data map[string]domain.Person
	mu   sync.RWMutex
}

func NewInMemoryPersonRepository() *InMemoryPersonRepository {
	return &InMemoryPersonRepository{
		data: make(map[string]domain.Person),
	}
}

func (r *InMemoryPersonRepository) GetAll() ([]domain.Person, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	persons := make([]domain.Person, 0, len(r.data))
	for _, person := range r.data {
		persons = append(persons, person)
	}
	return persons, nil
}

func (r *InMemoryPersonRepository) GetByID(id string) (domain.Person, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	person, exists := r.data[id]
	if !exists {
		return domain.Person{}, errors.New("person not found")
	}
	return person, nil
}

func (r *InMemoryPersonRepository) Create(person domain.Person) (domain.Person, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	person.ID = uuid.NewString()
	r.data[person.ID] = person
	return person, nil
}

func (r *InMemoryPersonRepository) Update(id string, person domain.Person) (domain.Person, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return domain.Person{}, errors.New("person not found")
	}
	person.ID = id
	r.data[id] = person
	return person, nil
}

func (r *InMemoryPersonRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return errors.New("person not found")
	}
	delete(r.data, id)
	return nil
}
