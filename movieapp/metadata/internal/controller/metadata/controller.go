package metadata

import (
	"context"
	"errors"

	"github.com/ullasasindhur/microservices_with_go/movieapp/metadata/internal/repository"
	model "github.com/ullasasindhur/microservices_with_go/movieapp/metadata/pkg"
)

// ErrNotFound is return when a requested record is not found.
var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller defines metadata service controller
type Controller struct {
	repo metadataRepository
}

// New creates a metadata service controller
func New(repo metadataRepository) *Controller {
	return &Controller{repo: repo}
}

// Get returns a movie metadata by id.
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
