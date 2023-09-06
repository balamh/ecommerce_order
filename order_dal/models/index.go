package models

// import("go.mongodb.org/mongo-driver/bson/primitive")

type Orders struct {
	// ID            primitive.ObjectID `json:"_id" bson:"_id"`
	CustomerId    string     `json:"customerid" bson:"customerid"`
	PaymentId     string     `json:"payment_id" bson:"payment_id" `
	PaymentStatus string     `json:"paymentstatus" bson:"paymentstatus"`
	Status        string     `json:"status" bson:"status"`
	Currency      string     `json:"currency" bson:"currency"`
	Items         []Items    `json:"items" bson:"items"`
	Shipping      []Shipping `json:"shipping" bson:"shipping"`
	Carrier       string     `json:"carrier" bson:"carrier"`
	Tracking      string     `json:"tracking" bson:"tracking"`
}

type Shipping struct {
	Address []Address `json:"address" bson:"address"`
	Origin  []Origin  `json:"origin" bson:"origin"`
}

type Items struct {
	Sku         string  `json:"sku" bson:"sku"`
	Quantity    string  `json:"quantity" bson:"quantity"`
	Price       float64 `json:"price" bson:"price"`
	Discount    float64 `json:"discount" bson:"discount"`
	PreTaxTotal float64 `json:"pretaxtotal" bson:"pretaxtotal"`
	Tax         float64 `json:"tax" bson:"tax"`
	Total       float64 `json:"total" bson:"total"`
}

type Address struct {
	Street1 string `json:"street1" bson:"street1"`
	Street2 string `json:"street2" bson:"street2"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Zip     string `json:"zip" bson:"zip"`
}
type Origin struct {
	Street1 string `json:"street1" bson:"street1"`
	Street2 string `json:"street2" bson:"street2"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Zip     string `json:"zip" bson:"zip"`
}
