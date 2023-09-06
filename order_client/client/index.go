package main

import (
	"context"
	// "fmt"

	"log"

	// models "ecommerce_order/order_dal/models"
	pb "ecommerce_order/order_proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:7010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)
	_, err1 := client.CreateOrder(context.Background(), &pb.CustomerOrder{
		CustomerId:    "33",
		PaymentId:     "your_payment_id",
		PaymentStatus: "your_payment_status",
		Status:        "your_status",
		Currency:      "your_currency",
		Items: []*pb.Items{
			{
				Sku:         "SKU001",
				Quantity:    "5",
				Price:       12.34, // Your price value
				Discount:    5.67,  // Your discount value
				PreTaxTotal: 18.01, // Your pre-tax total value
				Tax:         1.23,  // Your tax value
				Total:       19.24, // Your total value
			},
			// Add more items if needed
		},
		Shipping: []*pb.Shipping{
			{
				Address: []*pb.Address{
					{
						Street1: "Anna Nagar",
						Street2: "Gandhi nagar",
						City:    "Chennai",
						State:   "TamilNadu",
						Country: "India",
						Zip:     "56456",
					},
				},
				Origin: []*pb.Origin{
					{
						Street1: "Anna Nagar",
						Street2: "Gandhi nagar",
						City:    "Chennai",
						State:   "TamilNadu",
						Country: "India",
						Zip:     "56456",
					},
				},
			},
		},
		Carrier:  "your_carrier",
		Tracking: "your_tracking",
	})

	// _, err1 := client.RemoveOrderCustomer(context.Background(), &pb.RemoveOrderRequest{
	// 	CustomerId:345,
	// })

	// 	_, err1 := client.GetOrderDetails(context.Background(), &pb.GetOrderRequest{
	// 	CustomerId:76,
	// })
	if err1 != nil {
		log.Fatalf("Failed to call SayHello: %v", err1)
	}

}
