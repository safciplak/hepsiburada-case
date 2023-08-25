package products

import (
	"context"

	"github.com/hepsiburada/hepsiburada-command/app"
)

var (
	ErrorNotfound = app.BusinessError("campaign not found")
)

type products interface {
	CreateProduct(ctx context.Context, c *Product) error
	GetProduct(ctx context.Context, productName string) (*Product, error)
}

type Service struct {
	Products products
}

func (s *Service) CreateProduct(ctx context.Context, c *Product) error {
	return s.Products.CreateProduct(ctx, c)
}

func (s *Service) GetProduct(ctx context.Context, productName string) (*Product, error) {
	return s.Products.GetProduct(ctx, productName)
}
