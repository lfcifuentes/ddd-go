package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/lfcifuentes/ddd-go/aggregate"
	"github.com/lfcifuentes/ddd-go/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

// New is a factory function to generate a new repository of products
func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	products := make([]aggregate.Product, 0, len(mr.products))
	for _, product := range mr.products {
		products = append(products, product)
	}
	return products, nil
}

// GetByID searches for a product based on it's ID
func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

// Add will add a new product to the repository
func (mpr *MemoryProductRepository) Add(newprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	mpr.products[newprod.GetID()] = newprod

	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *MemoryProductRepository) Update(upprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

// Delete remove an product from the repository
func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
