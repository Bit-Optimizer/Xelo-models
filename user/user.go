package user

type Address struct {
	country  string `bson:"country"`
	division string `bson:"division"`
	city     string `bson:"city"`
	lane     string `bson:"lane"`
	pin      string `bson:"pin"`
}

type User struct {
	id      string  `bson:"_id,omitempty"`
	name    string  `bson:"name"`
	email   string  `bson:"email"`
	phone   string  `bson:"phone"`
	profile string  `bson:"profile"`
	address Address `bson:"address"`
}