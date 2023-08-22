package memory

import (
	"context"
	"sync"

	"github.com/ullasasindhur/microservices_with_go/movieapp/metadata/internal/repository"
	model "github.com/ullasasindhur/microservices_with_go/movieapp/metadata/pkg"
)

// Repository defines a memory movie metadata repository.
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new memory repository.
func New() *Repository {
	return &Repository{
		data: map[string]*model.Metadata{},
	}
}

// Get retrives movie metadata for by movie id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
