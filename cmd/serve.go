package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lfcifuentes/ddd-go/aggregate"
	"github.com/lfcifuentes/ddd-go/services/order"
	"github.com/lfcifuentes/ddd-go/services/tavern"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	Long:  ``,
	Run:   startServer,
}

func startServer(_ *cobra.Command, _ []string) {
	// create products array
	products := CreateProducts()

	fmt.Println("Products:")
	for _, p := range products {
		fmt.Printf("Name: %s, Description: %s, Price: %f\n", p.GetItem().Name, p.GetItem().Description, p.GetPrice())
	}

	// Create Order Service to use in tavern
	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(viper.GetString("MONGO_URI")),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		fmt.Println("Error creating order service")
		panic(err)
	}

	// Create tavern service
	tavern, err := tavern.NewTavern(tavern.WithOrderService(os))
	if err != nil {
		fmt.Println("Error creating tavern service")
		panic(err)
	}

	cus, err := os.AddCustomer("Percy")
	if err != nil {
		fmt.Println("Error creating customer")
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Order
	err = tavern.Order(cus, order)
	if err != nil {
		panic(err)
	}
}

func CreateProducts() []aggregate.Product {
	apple, err := aggregate.NewProduct("Apple", "Red apple", 0.5)
	if err != nil {
		panic(err)
	}

	banana, err := aggregate.NewProduct("Banana", "Yellow banana", 0.3)
	if err != nil {
		panic(err)
	}

	orage, err := aggregate.NewProduct("Orange", "Orange orange", 0.7)
	if err != nil {
		panic(err)
	}

	return []aggregate.Product{
		apple,
		banana,
		orage,
	}
}
