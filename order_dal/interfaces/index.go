// this interface will provide the requried methods
package interfaces

import (
	// "context"
	"ecommerce_order/order_dal/models"
	// ecommerce_order "ecommerce_order/order_proto"
)

type IOrder interface {
	CreateOrder(input *models.Orders) (models.Orders, error) 
	// RemoveOrder(Customer_ID string) (string, error)
	// GetAllOrder(CustomerId string) ([]models.Orders, error)
	// checkItemAvailability(ctx context.Context, sku string, quantity string) (bool, error)
}
