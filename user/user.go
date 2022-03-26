package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	country  string `bson:"country"`
	division string `bson:"division"`
	city     string `bson:"city"`
	lane     string `bson:"lane"`
	pin      string `bson:"pin"`
}

type User struct {
	id        string   `bson:"_id,omitempty"`
	name      string   `bson:"name"`
	email     string   `bson:"email"`
	phone     string   `bson:"phone"`
	profile   string   `bson:"profile"`
	address   Address  `bson:"address"`
	rating    float32  `bson:"rating"`
	ratecount int32    `bson:"ratecount"`
	wishlist  []string `bson:"wishlist"`
	products  []string `bson:"products"`
	createdAt primitive.DateTime `bson:"createdAt"`
	updatedAt primitive.DateTime `bson:"updatedAt"`
}