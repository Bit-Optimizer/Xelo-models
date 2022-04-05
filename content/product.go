package content

import "go.mongodb.org/mongo-driver/bson/primitive"

type Spec struct {
	Key			string 		`bson:"key"`
	Value		string		`bson:"value"`
}

type Address struct {
	Country  	string 		`bson:"country"`
	Division 	string 		`bson:"division"`
	City     	string 		`bson:"city"`
	Lane     	string 		`bson:"lane"`
	Pin      	string 		`bson:"pin"`
}

type Product struct {
	Id			string 		`bson:"_id,omitempty"`
	Price		int32 		`bson:"price"`
	Name		string		`bson:"name"`
	Description	string		`bson:"description"`
	Brand		string		`bson:"brand"`
	Category	string		`bson:"category"`
	Subcategory	string		`bson:"subcategory"`
	Address		Address 	`bson:"address"`
	Seller		string		`bson:"seller"`
	Status		string		`bson:"status"`
	UploadDate	primitive.DateTime `bson:"uploadDate"`
	Spec		[]Spec 		`bson:"spec"`
	Photourl	[]string	`bson:"photourl"`
	Mfdate 		primitive.DateTime `bson:"mfdate"`
	Likes		[]string 	`bson:"likes"`
}