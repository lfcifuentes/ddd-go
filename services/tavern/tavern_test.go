package tavern

import (
	"testing"

	"github.com/google/uuid"
	"github.com/lfcifuentes/ddd-go/aggregate"
	"github.com/lfcifuentes/ddd-go/services/order"
	"github.com/spf13/viper"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{
		beer, peenuts, wine,
	}
	return products
}
func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(viper.GetString("MONGO_URI_TEST")),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	uid, err := os.AddCustomer("LUIS")
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}

}

func Test_MongoTavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(viper.GetString("MONGO_URI_TEST")),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	uid, err := os.AddCustomer("LUIS")
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}

}
