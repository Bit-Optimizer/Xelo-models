package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	Country  string `bson:"country"`
	Division string `bson:"division"`
	City     string `bson:"city"`
	Lane     string `bson:"lane"`
	Pin      string `bson:"pin"`
}

type User struct {
	Id        string   `bson:"_id,omitempty"`
	Name      string   `bson:"name"`
	Email     string   `bson:"email"`
	Phone     string   `bson:"phone"`
	Profile   string   `bson:"profile"`
	Address   Address  `bson:"address"`
	Rating    float32  `bson:"rating"`
	Ratecount int32    `bson:"ratecount"`
	Wishlist  []string `bson:"wishlist"`
	Products  []string `bson:"products"`
	CreatedAt primitive.DateTime `bson:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt"`
}