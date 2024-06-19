package memory

import (
	"context"
	"sync"

	"github.com/an1l4/movieapp/metadata/internal/repository"
	"github.com/an1l4/movieapp/metadata/pkg"
)

// Repository defines a memory movie metadata repository
type Repository struct {
	sync.RWMutex
	data map[string]*pkg.Metadata
}

// New creates a new memory repository.
func New() *Repository {
	return &Repository{
		data: map[string]*pkg.Metadata{},
	}
}

func (r *Repository) Get(_ context.Context, id string) (*pkg.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

func (r *Repository) Put(_ context.Context, id string, metadata *pkg.Metadata) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = metadata
	return nil
}
